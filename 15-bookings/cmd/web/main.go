package main

import (
	"database/sql"
	"encoding/gob"
	"flag"
	"fmt"
	"github.com/alexedwards/scs/v2"
	"github.com/johnwr-response/golang-build-modern-web-applications/15-bookings/internal/config"
	"github.com/johnwr-response/golang-build-modern-web-applications/15-bookings/internal/driver"
	"github.com/johnwr-response/golang-build-modern-web-applications/15-bookings/internal/handlers"
	"github.com/johnwr-response/golang-build-modern-web-applications/15-bookings/internal/helpers"
	"github.com/johnwr-response/golang-build-modern-web-applications/15-bookings/internal/models"
	"github.com/johnwr-response/golang-build-modern-web-applications/15-bookings/internal/render"
	"log"
	"net/http"
	"os"
	"time"
)

const portNumber = ":8088"

var app config.AppConfig
var session *scs.SessionManager
var infoLog *log.Logger
var errorLog *log.Logger

// main is the main application function
func main() {

	db, err := run()
	if err != nil {
		log.Fatal(err)
	}
	defer func(SQL *sql.DB) {
		err := SQL.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(db.SQL)

	defer close(app.MailChan)

	fmt.Println("Starting sendmail listener...")
	listenForMail()

	fmt.Printf("Starting application on port: %s\n", portNumber)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)

}

func run() (*driver.DB, error) {
	// what am I going to put in the session
	gob.Register(models.Reservation{})
	gob.Register(models.User{})
	gob.Register(models.Room{})
	gob.Register(models.Restriction{})
	//gob.Register(models.RoomRestriction{})
	gob.Register(map[string]int{})

	// read flags
	inProduction := flag.Bool("production", false, "Application is in production mode")
	useCache := flag.Bool("cache", false, "Use template cache")
	dbName := flag.String("db-name", "bookings", "Database name")
	dbHost := flag.String("db-host", "localhost", "Database host")
	dbUser := flag.String("db-user", "postgres", "Database user")
	dbPass := flag.String("db-pass", "exampleDOT33", "Database password")
	dbPort := flag.String("db-port", "5433", "Database port")
	dbSSL := flag.String("db-ssl", "disable", "Database ssl settings (disable, prefer, require)")
	flag.Parse()

	if *dbName == "" || *dbHost == "" || *dbUser == "" || *dbPass == "" || *dbPort == "" {
		fmt.Println("Missing required flags")
		os.Exit(1)
	}

	// create channel for email
	mailChan := make(chan models.MailData)
	app.MailChan = mailChan

	// change this to true when in production
	app.InProduction = *inProduction
	app.UseCache = *useCache

	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	app.InfoLog = infoLog

	errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	app.ErrorLog = errorLog

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction
	app.Session = session

	// connect to database
	log.Println("Connecting to database...")
	connectionString := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=%s", *dbHost, *dbPort, *dbName, *dbUser, *dbPass, *dbSSL)
	db, err := driver.ConnectSQL(connectionString)
	//db, err := driver.ConnectSQL("host=localhost port=5433 dbname=bookings user=postgres password=exampleDOT33")
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}
	log.Println("Successfully connected to database")

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatalf("cannot create template cache: %v", err)
		return nil, err
	}
	app.TemplateCache = tc

	repo := handlers.NewRepo(&app, db)
	handlers.NewHandlers(repo)
	render.NewRenderer(&app)
	helpers.NewHelpers(&app)

	return db, nil
}

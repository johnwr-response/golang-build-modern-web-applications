package main

import (
	"fmt"
	"github.com/myorg/myapp/pkg/config"
	"github.com/myorg/myapp/pkg/handlers"
	"github.com/myorg/myapp/pkg/render"
	"log"
	"net/http"
)

const portNumber = ":8088"

// main is the main application function
func main() {
	var app config.AppConfig
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}
	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	fmt.Printf("Starting application on port: %s\n", portNumber)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)

}

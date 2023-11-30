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

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Printf("Starting application on port: %s\n", portNumber)
	err = http.ListenAndServe(portNumber, nil)
	if err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}

}

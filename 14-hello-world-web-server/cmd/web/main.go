package main

import (
	"fmt"
	"github.com/myorg/myapp/pkg/config"
	"github.com/myorg/myapp/pkg/handlers"
	"github.com/myorg/myapp/pkg/render"
	"log"
	"net/http"
)

const portNumber = ":8080"

// main is the main application function
func main() {
	var app config.AppConfig
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}
	app.TemplateCache = tc
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Starting application on port: %s ", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}

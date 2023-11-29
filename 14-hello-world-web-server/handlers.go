package main

import (
	"net/http"
)

// Home is the home page handler
func Home(w http.ResponseWriter, _ *http.Request) {
	renderTemplate(w, "home.page.tmpl")
}

// About is the about page handler
func About(w http.ResponseWriter, _ *http.Request) {
	renderTemplate(w, "about.page.tmpl")
}

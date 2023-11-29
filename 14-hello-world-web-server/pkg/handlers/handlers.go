package handlers

import (
	"github.com/myorg/myapp/pkg/render"
	"net/http"
)

// Home is the home page handler
func Home(w http.ResponseWriter, _ *http.Request) {
	render.RenderingTemplate(w, "home.page.tmpl")
}

// About is the about page handler
func About(w http.ResponseWriter, _ *http.Request) {
	render.RenderingTemplate(w, "about.page.tmpl")
}

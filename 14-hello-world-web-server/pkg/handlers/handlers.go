package handlers

import (
	"github.com/myorg/myapp/pkg/config"
	"github.com/myorg/myapp/pkg/models"
	"github.com/myorg/myapp/pkg/render"
	"net/http"
)

// Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository { return &Repository{App: a} }

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) { Repo = r }

// Home is the home page handler
func (m *Repository) Home(w http.ResponseWriter, _ *http.Request) {
	render.RenderingTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

// About is the about page handler
func (m *Repository) About(w http.ResponseWriter, _ *http.Request) {
	// perform some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again."

	// send the data to the template
	render.RenderingTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

package handlers

import (
	"github.com/johnwr-response/golang-build-modern-web-applications/15-bookings/pkg/config"
	"github.com/johnwr-response/golang-build-modern-web-applications/15-bookings/pkg/models"
	"github.com/johnwr-response/golang-build-modern-web-applications/15-bookings/pkg/render"
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
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderingTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

// About is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// perform some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again."
	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	// send the data to the template
	render.RenderingTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

// Contact renders the make a contact page and displays a form
func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	render.RenderingTemplate(w, "contact.page.tmpl", &models.TemplateData{})
}

// Reservation renders the make a reservation page and displays a form
func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	render.RenderingTemplate(w, "make-reservation.page.tmpl", &models.TemplateData{})
}

// Majors renders the generals-quarters room page
func (m *Repository) Generals(w http.ResponseWriter, r *http.Request) {
	render.RenderingTemplate(w, "generals.page.tmpl", &models.TemplateData{})
}

// Majors renders the majors-suite room page
func (m *Repository) Majors(w http.ResponseWriter, r *http.Request) {
	render.RenderingTemplate(w, "majors.page.tmpl", &models.TemplateData{})
}

// Availability renders the search availability page
func (m *Repository) Availability(w http.ResponseWriter, r *http.Request) {
	render.RenderingTemplate(w, "search-availability.page.tmpl", &models.TemplateData{})
}

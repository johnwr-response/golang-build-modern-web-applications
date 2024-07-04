package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/johnwr-response/golang-build-modern-web-applications/15-bookings/internal/config"
	"github.com/johnwr-response/golang-build-modern-web-applications/15-bookings/internal/forms"
	"github.com/johnwr-response/golang-build-modern-web-applications/15-bookings/internal/models"
	"github.com/johnwr-response/golang-build-modern-web-applications/15-bookings/internal/render"
	"log"
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

	render.RenderingTemplate(w, r, "home.page.tmpl", &models.TemplateData{})
}

// About is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// perform some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again."
	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	// send the data to the template
	render.RenderingTemplate(w, r, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

// Contact renders the make a contact page and displays a form
func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	render.RenderingTemplate(w, r, "contact.page.tmpl", &models.TemplateData{})
}

// Reservation renders the make a reservation page and displays a form
func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	var emptyReservation models.Reservation
	data := make(map[string]interface{})
	data["reservation"] = emptyReservation
	render.RenderingTemplate(w, r, "make-reservation.page.tmpl", &models.TemplateData{
		Form: forms.New(nil),
		Data: data,
	})
}

// PostReservation handles the posting of a reservation form
func (m *Repository) PostReservation(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		return
	}

	reservation := models.Reservation{
		FirstName: r.Form.Get("first_name"),
		LastName:  r.Form.Get("last_name"),
		Email:     r.Form.Get("email"),
		Phone:     r.Form.Get("phone"),
	}

	form := forms.New(r.PostForm)
	form.Required("first_name", "last_name", "email")
	form.MinLength("first_name", 3, r)
	form.IsEmail("email")
	if !form.Valid() {
		data := make(map[string]interface{})
		data["reservation"] = reservation
		render.RenderingTemplate(w, r, "make-reservation.page.tmpl", &models.TemplateData{
			Form: form,
			Data: data,
		})
		return
	}

}

// Generals renders the generals-quarters room page
func (m *Repository) Generals(w http.ResponseWriter, r *http.Request) {
	render.RenderingTemplate(w, r, "generals.page.tmpl", &models.TemplateData{})
}

// Majors renders the majors-suite room page
func (m *Repository) Majors(w http.ResponseWriter, r *http.Request) {
	render.RenderingTemplate(w, r, "majors.page.tmpl", &models.TemplateData{})
}

// Availability renders the search availability page
func (m *Repository) Availability(w http.ResponseWriter, r *http.Request) {
	render.RenderingTemplate(w, r, "search-availability.page.tmpl", &models.TemplateData{})
}

// PostAvailability renders the search availability page
func (m *Repository) PostAvailability(w http.ResponseWriter, r *http.Request) {
	//start2 := r.FormValue("start") ?!?
	start := r.Form.Get("start")
	end := r.Form.Get("end")
	_, err := w.Write([]byte(fmt.Sprintf("start date is %s and end date is %s", start, end)))
	if err != nil {
		return
	}
}

type jsonResponse struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
}

// AvailabilityJSON handles requests for availability and send JSON response
func (m *Repository) AvailabilityJSON(w http.ResponseWriter, r *http.Request) {
	resp := jsonResponse{
		OK:      true,
		Message: "Available!",
	}
	out, err := json.MarshalIndent(resp, "", "     ")
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(out)
	if err != nil {
		return
	}

}

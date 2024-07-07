package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/johnwr-response/golang-build-modern-web-applications/15-bookings/internal/config"
	"github.com/johnwr-response/golang-build-modern-web-applications/15-bookings/internal/driver"
	"github.com/johnwr-response/golang-build-modern-web-applications/15-bookings/internal/forms"
	"github.com/johnwr-response/golang-build-modern-web-applications/15-bookings/internal/helpers"
	"github.com/johnwr-response/golang-build-modern-web-applications/15-bookings/internal/models"
	"github.com/johnwr-response/golang-build-modern-web-applications/15-bookings/internal/render"
	"github.com/johnwr-response/golang-build-modern-web-applications/15-bookings/internal/repository"
	"github.com/johnwr-response/golang-build-modern-web-applications/15-bookings/internal/repository/dbrepo"
	"net/http"
	"strconv"
	"time"
)

// Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
	DB  repository.DatabaseRepo
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig, db *driver.DB) *Repository {
	return &Repository{App: a, DB: dbrepo.NewPostgresRepo(db.SQL, a)}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) { Repo = r }

// Home is the home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	err := render.Template(w, r, "home.page.tmpl", &models.TemplateData{})
	if err != nil {
		return
	}
}

// About is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// send the data to the template
	err := render.Template(w, r, "about.page.tmpl", &models.TemplateData{})
	if err != nil {
		return
	}
}

// Contact renders the make a contact page and displays a form
func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	err := render.Template(w, r, "contact.page.tmpl", &models.TemplateData{})
	if err != nil {
		return
	}
}

// Reservation renders the make a reservation page and displays a form
func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	var emptyReservation models.Reservation
	data := make(map[string]interface{})
	data["reservation"] = emptyReservation
	err := render.Template(w, r, "make-reservation.page.tmpl", &models.TemplateData{
		Form: forms.New(nil),
		Data: data,
	})
	if err != nil {
		return
	}
}

// PostReservation handles the posting of a reservation form
func (m *Repository) PostReservation(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	sd := r.Form.Get("start_date")
	ed := r.Form.Get("end_date")
	layout := "2006-01-02 15:04:05"
	startDate, err := time.Parse(layout, sd)
	if err != nil {
		helpers.ServerError(w, err)
	}
	endDate, err := time.Parse(layout, ed)
	if err != nil {
		helpers.ServerError(w, err)
	}
	roomID, err := strconv.Atoi(r.Form.Get("room_id"))
	if err != nil {
		helpers.ServerError(w, err)
	}

	reservation := models.Reservation{
		FirstName: r.Form.Get("first_name"),
		LastName:  r.Form.Get("last_name"),
		Email:     r.Form.Get("email"),
		Phone:     r.Form.Get("phone"),
		StartDate: startDate,
		EndDate:   endDate,
		RoomID:    roomID,
	}

	form := forms.New(r.PostForm)
	form.Required("first_name", "last_name", "email")
	form.MinLength("first_name", 3)
	form.IsEmail("email")
	if !form.Valid() {
		data := make(map[string]interface{})
		data["reservation"] = reservation
		err := render.Template(w, r, "make-reservation.page.tmpl", &models.TemplateData{
			Form: form,
			Data: data,
		})
		if err != nil {
			return
		}
		return
	}

	err = m.DB.InsertReservation(reservation)
	if err != nil {
		helpers.ServerError(w, err)
	}

	m.App.Session.Put(r.Context(), "reservation", reservation)

	http.Redirect(w, r, "/reservation-summary", http.StatusSeeOther)

}

// Generals renders the generals-quarters room page
func (m *Repository) Generals(w http.ResponseWriter, r *http.Request) {
	err := render.Template(w, r, "generals.page.tmpl", &models.TemplateData{})
	if err != nil {
		return
	}
}

// Majors renders the majors-suite room page
func (m *Repository) Majors(w http.ResponseWriter, r *http.Request) {
	err := render.Template(w, r, "majors.page.tmpl", &models.TemplateData{})
	if err != nil {
		return
	}
}

// Availability renders the search availability page
func (m *Repository) Availability(w http.ResponseWriter, r *http.Request) {
	err := render.Template(w, r, "search-availability.page.tmpl", &models.TemplateData{})
	if err != nil {
		return
	}
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
func (m *Repository) AvailabilityJSON(w http.ResponseWriter, _ *http.Request) {
	resp := jsonResponse{
		OK:      true,
		Message: "Available!",
	}
	out, err := json.MarshalIndent(resp, "", "     ")
	if err != nil {
		helpers.ServerError(w, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(out)
	if err != nil {
		return
	}

}

// ReservationSummary renders the reservation summary page
func (m *Repository) ReservationSummary(w http.ResponseWriter, r *http.Request) {
	reservation, ok := m.App.Session.Get(r.Context(), "reservation").(models.Reservation)
	if !ok {
		m.App.ErrorLog.Println("Cannot get error from session")
		m.App.Session.Put(r.Context(), "error", "Cannot get reservation from session")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	m.App.Session.Remove(r.Context(), "reservation")
	data := make(map[string]interface{})
	data["reservation"] = reservation
	err := render.Template(w, r, "reservation-summary.page.tmpl", &models.TemplateData{
		Data: data,
	})
	if err != nil {
		return
	}
}

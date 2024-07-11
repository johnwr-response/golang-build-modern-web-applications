package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/johnwr-response/golang-build-modern-web-applications/15-bookings/internal/driver"
	"github.com/johnwr-response/golang-build-modern-web-applications/15-bookings/internal/models"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"strings"
	"testing"
	"time"
)

var theTests = []struct {
	name               string
	url                string
	method             string
	expectedStatusCode int
}{
	{"home", "/", "GET", http.StatusOK},
	{"about", "/about", "GET", http.StatusOK},
	{"gq", "/generals-quarters", "GET", http.StatusOK},
	{"ms", "/majors-suite", "GET", http.StatusOK},
	{"sa", "/search-availability", "GET", http.StatusOK},
	{"contact", "/contact", "GET", http.StatusOK},
	{"non-existent", "/green/eggs/and/ham", "GET", http.StatusNotFound},
	{"login", "/user/login", "GET", http.StatusOK},
	{"logout", "/user/logout", "GET", http.StatusOK},
	{"dashboard", "/admin/dashboard", "GET", http.StatusOK},
	{"new res", "/admin/reservations-new", "GET", http.StatusOK},
	{"all res", "/admin/reservations-all", "GET", http.StatusOK},
	{"show res", "/admin/reservations/new/1/show", "GET", http.StatusOK},
}

func TestHandlers(t *testing.T) {
	routes := getRoutes()
	ts := httptest.NewTLSServer(routes)
	defer ts.Close()

	for _, e := range theTests {
		if e.method == "GET" {
			resp, err := ts.Client().Get(ts.URL + e.url)
			if err != nil {
				t.Log(err)
				t.Fatal(err)
			}
			if resp.StatusCode != e.expectedStatusCode {
				t.Errorf("for %s, expected %d but got %d", e.name, e.expectedStatusCode, resp.StatusCode)
			}
		}
	}
}

func TestRepository_Reservation(t *testing.T) {
	reservation := models.Reservation{
		RoomID: 1,
		Room: models.Room{
			ID:       1,
			RoomName: "General's Quarters",
		},
	}
	req, _ := http.NewRequest("GET", "/make-reservation", nil)
	ctx := getCtx(req)
	req = req.WithContext(ctx)

	rr := httptest.NewRecorder()
	session.Put(ctx, "reservation", reservation)
	handler := http.HandlerFunc(Repo.Reservation)
	handler.ServeHTTP(rr, req)
	if rr.Code != http.StatusOK {
		t.Errorf("Reservation handler returned wrong response code: got %v want %v", rr.Code, http.StatusOK)
	}

	// test case where reservation is not in session (reset everything)
	req, _ = http.NewRequest("GET", "/make-reservation", nil)
	ctx = getCtx(req)
	req = req.WithContext(ctx)
	rr = httptest.NewRecorder()
	handler.ServeHTTP(rr, req)
	if rr.Code != http.StatusSeeOther {
		t.Errorf("Reservation handler returned wrong response code: got %v want %v", rr.Code, http.StatusSeeOther)
	}

	// test case where reservation is a non-existent room
	req, _ = http.NewRequest("GET", "/make-reservation", nil)
	ctx = getCtx(req)
	req = req.WithContext(ctx)
	rr = httptest.NewRecorder()
	reservation.RoomID = 0
	session.Put(ctx, "reservation", reservation)
	handler.ServeHTTP(rr, req)
	if rr.Code != http.StatusSeeOther {
		t.Errorf("Reservation handler returned wrong response code: got %v want %v", rr.Code, http.StatusSeeOther)
	}

}

func TestRepository_PostReservation(t *testing.T) {
	format := "2006-01-02"
	sd, _ := time.Parse(format, "2050-01-01")
	ed, _ := time.Parse(format, "2050-01-02")
	reservation := models.Reservation{
		RoomID: 1,
		Room: models.Room{
			ID:       1,
			RoomName: "General's Quarters",
		},
		StartDate: sd,
		EndDate:   ed,
	}

	// basic test
	postedData := url.Values{}
	postedData.Add("first_name", "John")
	postedData.Add("last_name", "Smith")
	postedData.Add("email", "john@smith.com")
	postedData.Add("phone", "123456789")
	req, _ := http.NewRequest("POST", "/make-reservation", strings.NewReader(postedData.Encode()))
	ctx := getCtx(req)
	req = req.WithContext(ctx)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	rr := httptest.NewRecorder()
	session.Put(ctx, "reservation", reservation)
	handler := http.HandlerFunc(Repo.PostReservation)
	handler.ServeHTTP(rr, req)
	if rr.Code != http.StatusSeeOther {
		t.Errorf("PostReservation handler returned wrong response code: got %v want %v", rr.Code, http.StatusSeeOther)
	}

	// test for missing reservation in session
	req, _ = http.NewRequest("POST", "/make-reservation", strings.NewReader(postedData.Encode()))
	ctx = getCtx(req)
	req = req.WithContext(ctx)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	rr = httptest.NewRecorder()
	//handler = http.HandlerFunc(Repo.PostReservation)
	handler.ServeHTTP(rr, req)
	if rr.Code != http.StatusSeeOther {
		t.Errorf("PostReservation handler returned wrong response code: got %v want %v", rr.Code, http.StatusSeeOther)
	}

	// test for missing POST body
	req, _ = http.NewRequest("POST", "/make-reservation", nil)
	ctx = getCtx(req)
	req = req.WithContext(ctx)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	rr = httptest.NewRecorder()
	session.Put(ctx, "reservation", reservation)
	//handler = http.HandlerFunc(Repo.PostReservation)
	handler.ServeHTTP(rr, req)
	if rr.Code != http.StatusSeeOther {
		t.Errorf("PostReservation handler returned wrong response code for missing post body: got %v want %v", rr.Code, http.StatusSeeOther)
	}

	// test for invalid data
	postedData = url.Values{}
	postedData.Add("first_name", "J")
	postedData.Add("last_name", "Smith")
	postedData.Add("email", "john@smith.com")
	postedData.Add("phone", "123456789")

	req, _ = http.NewRequest("POST", "/make-reservation", strings.NewReader(postedData.Encode()))
	ctx = getCtx(req)
	req = req.WithContext(ctx)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	rr = httptest.NewRecorder()
	session.Put(ctx, "reservation", reservation)
	//handler = http.HandlerFunc(Repo.PostReservation)
	handler.ServeHTTP(rr, req)
	if rr.Code != http.StatusSeeOther {
		t.Errorf("PostReservation handler returned wrong response code for validation of first_name in post body: got %v want %v", rr.Code, http.StatusSeeOther)
	}

	// test for failure to insert reservation into database
	postedData = url.Values{}
	postedData.Add("first_name", "Invalid")
	postedData.Add("last_name", "Smith")
	postedData.Add("email", "john@smith.com")
	postedData.Add("phone", "123456789")
	req, _ = http.NewRequest("POST", "/make-reservation", strings.NewReader(postedData.Encode()))
	ctx = getCtx(req)
	req = req.WithContext(ctx)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	rr = httptest.NewRecorder()
	session.Put(ctx, "reservation", reservation)
	//handler = http.HandlerFunc(Repo.PostReservation)
	handler.ServeHTTP(rr, req)
	if rr.Code != http.StatusSeeOther {
		t.Errorf("PostReservation handler failed when trying to fail inserting reservation, returned wrong response code for validation of first_name in post body: got %v want %v", rr.Code, http.StatusSeeOther)
	}

	// test for failure to insert restriction into database
	postedData = url.Values{}
	postedData.Add("first_name", "John")
	postedData.Add("last_name", "Smith")
	postedData.Add("email", "john@smith.com")
	postedData.Add("phone", "123456789")
	req, _ = http.NewRequest("POST", "/make-reservation", strings.NewReader(postedData.Encode()))
	ctx = getCtx(req)
	req = req.WithContext(ctx)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	rr = httptest.NewRecorder()
	reservation.RoomID = -1
	session.Put(ctx, "reservation", reservation)
	//handler = http.HandlerFunc(Repo.PostReservation)
	handler.ServeHTTP(rr, req)
	if rr.Code != http.StatusSeeOther {
		t.Errorf("PostReservation handler failed when trying to fail inserting reservation, returned wrong response code for validation of first_name in post body: got %v want %v", rr.Code, http.StatusSeeOther)
	}
}

func TestNewRepo(t *testing.T) {
	var db driver.DB
	testRepo := NewRepo(&app, &db)
	if reflect.TypeOf(testRepo).String() != "*handlers.Repository" {
		t.Errorf("Did not get correct type from NewRepo: got %v wanted *Repository", reflect.TypeOf(testRepo).String())
	}
}

func TestRepository_PostAvailability(t *testing.T) {
	// rooms are available
	postedData := url.Values{}
	postedData.Add("start", "2040-01-01")
	postedData.Add("end", "2040-01-02")
	req, _ := http.NewRequest("POST", "/search-availability", strings.NewReader(postedData.Encode()))
	ctx := getCtx(req)
	req = req.WithContext(ctx)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Repo.PostAvailability)
	handler.ServeHTTP(rr, req)
	if rr.Code != http.StatusOK {
		t.Errorf("Post availability when rooms are available gave wrong status code: got %v wanted %v", rr.Code, http.StatusOK)
	}

	// rooms are not available
	postedData = url.Values{}
	postedData.Add("start", "2050-01-01")
	postedData.Add("end", "2050-01-02")
	req, _ = http.NewRequest("POST", "/search-availability", strings.NewReader(postedData.Encode()))
	ctx = getCtx(req)
	req = req.WithContext(ctx)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	rr = httptest.NewRecorder()
	handler.ServeHTTP(rr, req)
	if rr.Code != http.StatusSeeOther {
		t.Errorf("Post availability when rooms are not available gave wrong status code: got %v wanted %v", rr.Code, http.StatusSeeOther)
	}

	// empty post body
	req, _ = http.NewRequest("POST", "/search-availability", nil)
	ctx = getCtx(req)
	req = req.WithContext(ctx)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	rr = httptest.NewRecorder()
	handler.ServeHTTP(rr, req)
	if rr.Code != http.StatusSeeOther {
		t.Errorf("Post availability with empty request body gave wrong status code: got %v wanted %v", rr.Code, http.StatusSeeOther)
	}

	// wrong format start date
	postedData = url.Values{}
	postedData.Add("start", "invalid")
	postedData.Add("end", "2040-01-02")
	req, _ = http.NewRequest("POST", "/search-availability", strings.NewReader(postedData.Encode()))
	ctx = getCtx(req)
	req = req.WithContext(ctx)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	rr = httptest.NewRecorder()
	handler.ServeHTTP(rr, req)
	if rr.Code != http.StatusSeeOther {
		t.Errorf("Post availability with invalid start date gave wrong status code: got %v wanted %v", rr.Code, http.StatusSeeOther)
	}

	// wrong format end date
	postedData = url.Values{}
	postedData.Add("start", "2040-01-01")
	postedData.Add("end", "invalid")
	req, _ = http.NewRequest("POST", "/search-availability", strings.NewReader(postedData.Encode()))
	ctx = getCtx(req)
	req = req.WithContext(ctx)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	rr = httptest.NewRecorder()
	handler.ServeHTTP(rr, req)
	if rr.Code != http.StatusSeeOther {
		t.Errorf("Post availability with invalid end date gave wrong status code: got %v wanted %v", rr.Code, http.StatusSeeOther)
	}

	// database query fails
	postedData = url.Values{}
	postedData.Add("start", "2060-01-01")
	postedData.Add("end", "2060-01-02")
	req, _ = http.NewRequest("POST", "/search-availability", strings.NewReader(postedData.Encode()))
	ctx = getCtx(req)
	req = req.WithContext(ctx)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	rr = httptest.NewRecorder()
	handler.ServeHTTP(rr, req)
	if rr.Code != http.StatusSeeOther {
		t.Errorf("Post availability when database query fails gave wrong status code: got %v wanted %v", rr.Code, http.StatusSeeOther)
	}

}

func TestRepository_AvailabilityJSON(t *testing.T) {
	// first case is rooms are not available
	postedData := url.Values{}
	postedData.Add("start", "2050-01-01")
	postedData.Add("end", "2050-01-02")
	postedData.Add("room_id", "1")

	req, _ := http.NewRequest("POST", "/search-availability-json", strings.NewReader(postedData.Encode()))
	ctx := getCtx(req)
	req = req.WithContext(ctx)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Repo.AvailabilityJSON)
	handler.ServeHTTP(rr, req)
	var j jsonResponse
	err := json.Unmarshal(rr.Body.Bytes(), &j)
	//t.Logf("rr.body : %v", rr.Body)
	if err != nil {
		t.Error("Failed to parse JSON")
	}
	if j.OK {
		t.Error("Got availability when none expected in AvailabilityJSON")
	}

	// case rooms are available
	postedData = url.Values{}
	postedData.Add("start", "2030-01-01")
	postedData.Add("end", "2030-01-02")
	postedData.Add("room_id", "1")

	req, _ = http.NewRequest("POST", "/search-availability-json", strings.NewReader(postedData.Encode()))
	ctx = getCtx(req)
	req = req.WithContext(ctx)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	rr = httptest.NewRecorder()
	handler.ServeHTTP(rr, req)
	j = jsonResponse{}
	err = json.Unmarshal(rr.Body.Bytes(), &j)
	//t.Logf("rr.body : %v", rr.Body)
	if err != nil {
		t.Error("Failed to parse JSON")
	}
	if !j.OK {
		t.Error("Got no availability when expected in AvailabilityJSON")
	}

	// no request body
	req, _ = http.NewRequest("POST", "/search-availability-json", nil)
	ctx = getCtx(req)
	req = req.WithContext(ctx)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	rr = httptest.NewRecorder()
	handler.ServeHTTP(rr, req)
	j = jsonResponse{}
	err = json.Unmarshal(rr.Body.Bytes(), &j)
	//t.Logf("rr.body : %v", rr.Body)
	if err != nil {
		t.Error("Failed to parse JSON")
	}
	if j.OK || j.Message != "Internal server error" {
		t.Error("Got availability when request body was empty")
	}

	// database error
	postedData = url.Values{}
	postedData.Add("start", "2060-01-01")
	postedData.Add("end", "2060-01-02")
	postedData.Add("room_id", "1")
	req, _ = http.NewRequest("POST", "/search-availability-json", strings.NewReader(postedData.Encode()))
	ctx = getCtx(req)
	req = req.WithContext(ctx)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	rr = httptest.NewRecorder()
	handler.ServeHTTP(rr, req)
	j = jsonResponse{}
	err = json.Unmarshal(rr.Body.Bytes(), &j)
	//t.Logf("rr.body : %v", rr.Body)
	if err != nil {
		t.Error("Failed to parse JSON")
	}
	if j.OK || j.Message != "Error connecting to database" {
		t.Error("Got no availability when expected in AvailabilityJSON")
	}

}

func TestRepository_ReservationSummary(t *testing.T) {
	// reservation in session
	reservation := models.Reservation{
		RoomID: 1,
		Room: models.Room{
			ID:       1,
			RoomName: "General's Quarters",
		},
	}
	req, _ := http.NewRequest("GET", "/reservation-summary", nil)
	ctx := getCtx(req)
	req = req.WithContext(ctx)
	rr := httptest.NewRecorder()
	session.Put(ctx, "reservation", reservation)
	handler := http.HandlerFunc(Repo.ReservationSummary)
	handler.ServeHTTP(rr, req)
	if rr.Code != http.StatusOK {
		t.Errorf("ReservationSummary handler returned wrong response code: got %v wanted %v", rr.Code, http.StatusOK)
	}

	// reservation not in session
	req, _ = http.NewRequest("GET", "/reservation-summary", nil)
	ctx = getCtx(req)
	req = req.WithContext(ctx)
	rr = httptest.NewRecorder()
	handler.ServeHTTP(rr, req)
	if rr.Code != http.StatusSeeOther {
		t.Errorf("ReservationSummary handler returned wrong response code: got %v wanted %v", rr.Code, http.StatusSeeOther)
	}
}

func TestRepository_ChooseRoom(t *testing.T) {
	// reservation in session
	reservation := models.Reservation{
		RoomID: 1,
		Room: models.Room{
			ID:       1,
			RoomName: "General's Quarters",
		},
	}
	req, _ := http.NewRequest("GET", "/choose-room/1", nil)
	ctx := getCtx(req)
	req = req.WithContext(ctx)
	req.RequestURI = "/choose-room/1"
	rr := httptest.NewRecorder()
	session.Put(ctx, "reservation", reservation)
	handler := http.HandlerFunc(Repo.ChooseRoom)
	handler.ServeHTTP(rr, req)
	if rr.Code != http.StatusSeeOther {
		t.Errorf("ChooseRoom handler returned wrong response code: got %v wanted %v", rr.Code, http.StatusSeeOther)
	}

	// reservation not in session
	req, _ = http.NewRequest("GET", "/choose-room/1", nil)
	ctx = getCtx(req)
	req = req.WithContext(ctx)
	req.RequestURI = "/choose-room/1"
	rr = httptest.NewRecorder()
	handler.ServeHTTP(rr, req)
	if rr.Code != http.StatusSeeOther {
		t.Errorf("ChooseRoom handler returned wrong response code: got %v wanted %v", rr.Code, http.StatusSeeOther)
	}

	// missing or malformed url parameter
	req, _ = http.NewRequest("GET", "/choose-room/wrong", nil)
	ctx = getCtx(req)
	req = req.WithContext(ctx)
	req.RequestURI = "/choose-room/wrong"
	rr = httptest.NewRecorder()
	handler.ServeHTTP(rr, req)
	if rr.Code != http.StatusSeeOther {
		t.Errorf("ChooseRoom handler returned wrong response code: got %v wanted %v", rr.Code, http.StatusSeeOther)
	}
}

func TestRepository_BookRoom(t *testing.T) {
	reservation := models.Reservation{
		RoomID: 1,
		Room: models.Room{
			ID:       1,
			RoomName: "General's Quarters",
		},
	}
	// database works
	req, _ := http.NewRequest("GET", "/book-room?s=2050-01-01&e2050-01-02&id=1", nil)
	ctx := getCtx(req)
	req = req.WithContext(ctx)
	rr := httptest.NewRecorder()
	session.Put(ctx, "reservation", reservation)
	handler := http.HandlerFunc(Repo.BookRoom)
	handler.ServeHTTP(rr, req)
	if rr.Code != http.StatusSeeOther {
		t.Errorf("BookRoom handler returned wrong response code: got %v wanted %v", rr.Code, http.StatusSeeOther)
	}
	// database failed
	req, _ = http.NewRequest("GET", "/book-room?s=2050-01-01&e2050-01-02&id=0", nil)
	ctx = getCtx(req)
	req = req.WithContext(ctx)
	rr = httptest.NewRecorder()
	session.Put(ctx, "reservation", reservation)
	handler.ServeHTTP(rr, req)
	if rr.Code != http.StatusSeeOther {
		t.Errorf("BookRoom handler returned wrong response code: got %v wanted %v", rr.Code, http.StatusSeeOther)
	}
}

// loginTests is the data for the Login handler tests
var loginTests = []struct {
	name               string
	email              string
	expectedStatusCode int
	expectedHTML       string
	expectedLocation   string
}{
	{"valid-credentials", "me@here.ca", http.StatusSeeOther, "", "/"},
	{"invalid-credentials", "jack@nimble.com", http.StatusSeeOther, "", "/user/login"},
	{"invalid-data", "j", http.StatusOK, `action="/user/login"`, ""},
}

// TestLogin is the Login handler tests
func TestLogin(t *testing.T) {
	// range through all tests
	for _, e := range loginTests {
		postedData := url.Values{}
		postedData.Add("email", e.email)
		postedData.Add("password", "password")

		// create request
		req, _ := http.NewRequest("POST", "/user/login", strings.NewReader(postedData.Encode()))
		ctx := getCtx(req)
		req = req.WithContext(ctx)

		// set the header
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		rr := httptest.NewRecorder()

		// call the handler
		handler := http.HandlerFunc(Repo.PostShowLogin)
		handler.ServeHTTP(rr, req)

		// new do some tests
		if rr.Code != e.expectedStatusCode {
			t.Errorf("failed %s: expected code %d, but got %d", e.name, e.expectedStatusCode, rr.Code)
		}
		if e.expectedLocation != "" {
			actualLoc, _ := rr.Result().Location()
			if actualLoc.String() != e.expectedLocation {
				t.Errorf("failed %s: expected location %s, but got %s", e.name, e.expectedLocation, actualLoc.String())
			}
		}
		if e.expectedHTML != "" {
			html := rr.Body.String()
			if !strings.Contains(html, e.expectedHTML) {
				t.Errorf("failed %s: expected html %s, but got %s", e.name, e.expectedHTML, html)
			}
		}

	}
}

// adminPostShowReservationTests is the data for the AdminPostShowReservation handler
var adminPostShowReservationTests = []struct {
	name               string
	url                string
	expectedStatusCode int
	expectedLocation   string
	expectedHTML       string
	postedData         url.Values
}{
	{name: "valid-data-from-new", url: "/admin/reservations/new/1/show", expectedStatusCode: http.StatusSeeOther, expectedLocation: "/admin/reservations-new", expectedHTML: "", postedData: url.Values{
		"first_name": {"John"}, "last_name": {"Smith"}, "email": {"john@smith.com"}, "phone": {"555-555-55555"},
	}},
	{name: "valid-data-from-all", url: "/admin/reservations/all/1/show", expectedStatusCode: http.StatusSeeOther, expectedLocation: "/admin/reservations-all", expectedHTML: "", postedData: url.Values{
		"first_name": {"John"}, "last_name": {"Smith"}, "email": {"john@smith.com"}, "phone": {"555-555-55555"},
	}},
	{name: "valid-data-from-cal", url: "/admin/reservations/cal/1/show", expectedStatusCode: http.StatusSeeOther, expectedLocation: "/admin/reservations-calendar?y=2024&m=07", expectedHTML: "", postedData: url.Values{
		"first_name": {"John"}, "last_name": {"Smith"}, "email": {"john@smith.com"}, "phone": {"555-555-55555"}, "year": {"2024"}, "month": {"07"},
	}},
}

// TestAdminPostShowReservation the AdminPostShowReservation handler
func TestAdminPostShowReservation(t *testing.T) {
	for _, e := range adminPostShowReservationTests {
		var req *http.Request
		if e.postedData != nil {
			req, _ = http.NewRequest("POST", "/admin/user/login", strings.NewReader(e.postedData.Encode()))
		} else {
			req, _ = http.NewRequest("POST", "/admin/user/login", nil)
		}
		ctx := getCtx(req)
		req = req.WithContext(ctx)
		//	req.RequestURI = e.url
		//	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		//	rr := httptest.NewRecorder()
		//	handler := http.HandlerFunc(Repo.AdminPostShowReservation)
		//	handler.ServeHTTP(rr, req)
		//	if rr.Code != e.expectedStatusCode {
		//		t.Errorf("failed %s: expected code %d, but got %d", e.name, e.expectedStatusCode, rr.Code)
		//	}
		//	if e.expectedLocation != "" {
		//		// get the URL from test
		//		actualLoc, _ := rr.Result().Location()
		//		if actualLoc.String() != e.expectedLocation {
		//			t.Errorf("failed %s: expected location %s, but got %s", e.name, e.expectedLocation, actualLoc.String())
		//		}
		//	}
		//	if e.expectedHTML != "" {
		//		// read the response body into a string
		//		html := rr.Body.String()
		//		if !strings.Contains(html, e.expectedHTML) {
		//			t.Errorf("failed %s: expected html %s, but got %s", e.name, e.expectedHTML, html)
		//		}
		//	}
	}
}

// adminPostReservationCalendarTests is the data for the PostReservationCalendar handler
var adminPostReservationCalendarTests = []struct {
	name               string
	postedData         url.Values
	expectedStatusCode int
	expectedLocation   string
	expectedHTML       string
	blocks             int
	reservations       int
}{
	{name: "cal", expectedStatusCode: http.StatusSeeOther, postedData: url.Values{
		"year":  {time.Now().Format("2006")},
		"month": {time.Now().Format("01")},
		fmt.Sprintf("add_block_1_%s", time.Now().AddDate(0, 0, 2).Format("2006-01-2")): {"1"},
	}},
	{name: "cal-blocks", expectedStatusCode: http.StatusSeeOther, blocks: 1, postedData: url.Values{}},
	{name: "cal-res", expectedStatusCode: http.StatusSeeOther, reservations: 1, postedData: url.Values{}},
}

// TestPostReservationCalendar tests the PostReservationCalendar handler
func TestPostReservationCalendar(t *testing.T) {
	for _, e := range adminPostReservationCalendarTests {
		var req *http.Request
		if e.postedData != nil {
			req, _ = http.NewRequest("POST", "/admin/reservations-calendar", strings.NewReader(e.postedData.Encode()))
		} else {
			req, _ = http.NewRequest("POST", "/admin/reservations-calendar", nil)
		}
		ctx := getCtx(req)
		req = req.WithContext(ctx)
		//	now := time.Now()
		//	bm := make(map[string]int)
		//	rm := make(map[string]int)
		//	currentYear, currentMonth, _ := now.Date()
		//	currentLocation := now.Location()
		//	firstOfMonth := time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, currentLocation)
		//	lastOfMonth := firstOfMonth.AddDate(0, 1, -1)
		//	for d := firstOfMonth; d.After(lastOfMonth) == false; d = d.AddDate(0, 0, 1) {
		//		rm[d.Format("2006-01-2")] = 0
		//		bm[d.Format("2006-01-2")] = 0
		//	}
		//	if e.blocks > 0 {
		//		bm[firstOfMonth.Format("2006-01-2")] = e.blocks
		//	}
		//	if e.reservations > 0 {
		//		rm[lastOfMonth.Format("2006-01-2")] = e.reservations
		//	}
		//	session.Put(ctx, "block_map_1", bm)
		//	session.Put(ctx, "reservation_map_1", rm)
		//	req.Header.Set("content-type", "application/x-www-form-urlencoded")
		//	rr := httptest.NewRecorder()
		//	handler := http.HandlerFunc(Repo.AdminPostReservationsCalendar)
		//	handler.ServeHTTP(rr, req)
		//
		//	if rr.Code != e.expectedStatusCode {
		//		t.Errorf("failed %s: expected code %d, but got %d", e.name, e.expectedStatusCode, rr.Code)
		//	}
		//
	}
}

// adminProcessReservationTests is the data for the AdminProcessReservation handler
var adminProcessReservationTests = []struct {
	name               string
	queryParams        string
	expectedStatusCode int
	expectedLocation   string
}{
	{name: "process-reservation", queryParams: "", expectedStatusCode: http.StatusSeeOther, expectedLocation: ""},
	{name: "process-reservation-back-to-cal", queryParams: "?y=2024&m=07", expectedStatusCode: http.StatusSeeOther, expectedLocation: ""},
}

// TestAdminProcessReservation tests the AdminProcessReservation handler
func TestAdminProcessReservation(t *testing.T) {
	for _, e := range adminProcessReservationTests {
		req, _ := http.NewRequest("GET", fmt.Sprintf("/admin/process-reservation/cal/1/do%s", e.queryParams), nil)
		ctx := getCtx(req)
		req = req.WithContext(ctx)
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(Repo.AdminProcessReservation)
		handler.ServeHTTP(rr, req)
		if rr.Code != e.expectedStatusCode {
			t.Errorf("failed %s: expected code %d, but got %d", e.name, e.expectedStatusCode, rr.Code)
		}
	}
}

// adminDeleteReservationTests is the data for the AdminDeleteReservation handler tests
var adminDeleteReservationTests = []struct {
	name               string
	queryParams        string
	expectedStatusCode int
	expectedLocation   string
}{
	{name: "delete-reservation", queryParams: "", expectedStatusCode: http.StatusSeeOther, expectedLocation: ""},
	{name: "delete-reservation-back-to-cal", queryParams: "?y=2024&m=07", expectedStatusCode: http.StatusSeeOther, expectedLocation: ""},
}

// TestAdminDeleteReservation tests the AdminDeleteReservation handler
func TestAdminDeleteReservation(t *testing.T) {
	for _, e := range adminDeleteReservationTests {
		req, _ := http.NewRequest("GET", fmt.Sprintf("/admin/delete-reservation/cal/1/do%s", e.queryParams), nil)
		ctx := getCtx(req)
		req = req.WithContext(ctx)
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(Repo.AdminDeleteReservation)
		handler.ServeHTTP(rr, req)
		if rr.Code != e.expectedStatusCode {
			t.Errorf("failed %s: expected code %d, but got %d", e.name, e.expectedStatusCode, rr.Code)
		}
	}
}

func getCtx(req *http.Request) context.Context {
	ctx, err := session.Load(req.Context(), req.Header.Get("X-Session"))
	if err != nil {
		log.Println(err)
	}
	return ctx
}

package repository

import (
	"github.com/johnwr-response/golang-build-modern-web-applications/15-bookings/internal/models"
	"time"
)

type DatabaseRepo interface {
	AllUsers() bool
	InsertReservation(res models.Reservation) (int, error)
	InsertRoomRestriction(rr models.RoomRestriction) error
	SearchAvailabilityByDates(start, end time.Time, roomID int)
}

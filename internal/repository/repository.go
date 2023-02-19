package repository

import (
	"time"

	"github.com/fxfrancky/bookings/internal/models"
)

type DatabaseRepo interface {
	AllUsers() bool

	InsertReservation(res models.Reservation) (int, error)
	InsertRoomRestriction(r models.RoomRestriction) error
	SearchAvailabilitiesByDatesByRoomID(start, end time.Time, roomID int) (bool, error)
	SearchAvailabilitiesForAllRooms(start, end time.Time) ([]models.Room, error)
	GetRoomByID(id int) (models.Room, error)
}

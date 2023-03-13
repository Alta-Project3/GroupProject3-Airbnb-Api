package reservations

import (
	"groupproject3-airbnb-api/features/rooms"
	"groupproject3-airbnb-api/features/user"
)

type ReservationEntity struct {
	Id                uint
	UserId            uint
	User              user.Core
	RoomId            uint
	Room              rooms.RoomEntity
	DateStart         string
	DateEnd           string
	Duration          int
	TotalPrice        int
	CardName          string
	CardNumber        string
	CardCvv           string
	CardMonth         string
	CardYear          string
	StatusReservation string
}

type ReservationServiceInterface interface {
	GetByRoomAndDateRange() ([]ReservationEntity, error)
	Create(reservationEntity ReservationEntity) (ReservationEntity, error)
}

type ReservationDataInterface interface {
	SelectyRoomAndDateRange() ([]ReservationEntity, error)
	Store(reservationEntity ReservationEntity) (uint, error)
}

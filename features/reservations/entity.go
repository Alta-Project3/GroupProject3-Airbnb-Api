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

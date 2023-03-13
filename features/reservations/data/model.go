package data

import (
	"groupproject3-airbnb-api/features/reservations"
	room "groupproject3-airbnb-api/features/rooms"
	roomData "groupproject3-airbnb-api/features/rooms/data"
	user "groupproject3-airbnb-api/features/user"
	userData "groupproject3-airbnb-api/features/user/data"
	"reflect"

	"gorm.io/gorm"
)

type Reservation struct {
	gorm.Model
	UserId            uint
	User              userData.User `gorm:"foreignKey:UserId"`
	RoomId            uint
	Room              roomData.Room `gorm:"foreignKey:RoomId"`
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

func ReservationEntityToReservation(reservationEntity reservations.ReservationEntity) Reservation {
	return Reservation{
		UserId:            reservationEntity.UserId,
		RoomId:            reservationEntity.RoomId,
		DateStart:         reservationEntity.DateStart,
		DateEnd:           reservationEntity.DateEnd,
		Duration:          reservationEntity.Duration,
		TotalPrice:        reservationEntity.TotalPrice,
		CardName:          reservationEntity.CardName,
		CardNumber:        reservationEntity.CardNumber,
		CardCvv:           reservationEntity.CardCvv,
		CardMonth:         reservationEntity.CardMonth,
		CardYear:          reservationEntity.CardYear,
		StatusReservation: reservationEntity.StatusReservation,
	}
}

func ReservationToReservationEntity(reservation Reservation) reservations.ReservationEntity {
	result := reservations.ReservationEntity{
		Id:                reservation.ID,
		UserId:            reservation.UserId,
		RoomId:            reservation.RoomId,
		DateStart:         reservation.DateStart,
		DateEnd:           reservation.DateEnd,
		Duration:          reservation.Duration,
		TotalPrice:        reservation.TotalPrice,
		CardName:          reservation.CardName,
		CardNumber:        reservation.CardNumber,
		CardCvv:           reservation.CardCvv,
		CardMonth:         reservation.CardMonth,
		CardYear:          reservation.CardYear,
		StatusReservation: reservation.StatusReservation,
	}

	if !reflect.ValueOf(reservation.User).IsZero() {
		result.User = user.Core{
			Name: reservation.User.Name,
		}
	}

	if !reflect.ValueOf(reservation.Room).IsZero() {
		result.Room = room.RoomEntity{
			RoomName: reservation.Room.RoomName,
		}
	}
	return result
}

func ListReservationToReservationEntity(reservation []Reservation) []reservations.ReservationEntity {
	var reservationsEntity []reservations.ReservationEntity
	for _, v := range reservation {
		reservationsEntity = append(reservationsEntity, ReservationToReservationEntity(v))
	}
	return reservationsEntity
}

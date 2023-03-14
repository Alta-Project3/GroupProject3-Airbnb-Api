package data

import (
	"groupproject3-airbnb-api/features/reservations"

	"gorm.io/gorm"
)

type query struct {
	db *gorm.DB
}

func New(db *gorm.DB) reservations.ReservationDataInterface {
	return &query{
		db: db,
	}
}

func (q *query) SelectyRoomAndDateRange(reservationEntity reservations.ReservationEntity) ([]reservations.ReservationEntity, error) {
	var reservations []Reservation
	err := q.db.Preload("User").
		Where("room_id = ? AND (date_start BETWEEN ? AND ? OR date_end BETWEEN ? AND ?)", reservationEntity.RoomId, reservationEntity.DateStart, reservationEntity.DateEnd, reservationEntity.DateStart, reservationEntity.DateEnd).
		Find(&reservations)
	if err.Error != nil {
		return nil, err.Error
	}
	return ListReservationToReservationEntity(reservations), nil
}

// Store implements reservations.ReservationDataInterface
func (*query) Store(reservationEntity reservations.ReservationEntity) (uint, error) {
	panic("unimplemented")
}

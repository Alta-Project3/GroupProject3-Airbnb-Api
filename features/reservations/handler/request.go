package handler

import "groupproject3-airbnb-api/features/reservations"

type ReservationRequest struct {
	RoomId    uint   `json:"room_id" form:"room_id"`
	DateStart string `json:"date_start" form:"date_start"`
	DateEnd   string `json:"date_end" form:"date_end"`
}

func ReservationRequestToReservationEntity(reservationRequest *ReservationRequest) reservations.ReservationEntity {
	return reservations.ReservationEntity{
		RoomId:    reservationRequest.RoomId,
		DateStart: reservationRequest.DateStart,
		DateEnd:   reservationRequest.DateEnd,
	}
}

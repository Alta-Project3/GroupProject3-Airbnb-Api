package handler

import "groupproject3-airbnb-api/features/reservations"

type ReservationRequest struct {
	RoomId     uint   `json:"room_id" form:"room_id"`
	DateStart  string `json:"date_start" form:"date_start"`
	DateEnd    string `json:"date_end" form:"date_end"`
	CardName   string `json:"card_name" form:"card_name"`
	CardNumber string `json:"card_number" form:"card_number"`
	CardCvv    string `json:"card_cvv" form:"card_cvv"`
	CardMonth  string `json:"card_month" form:"card_month"`
	CardYear   string `json:"card_yaer" form:"card_yaer"`
}

func ReservationRequestToReservationEntity(reservationRequest *ReservationRequest) reservations.ReservationEntity {
	return reservations.ReservationEntity{
		RoomId:     reservationRequest.RoomId,
		DateStart:  reservationRequest.DateStart,
		DateEnd:    reservationRequest.DateEnd,
		CardName:   reservationRequest.CardName,
		CardNumber: reservationRequest.CardNumber,
		CardCvv:    reservationRequest.CardCvv,
		CardMonth:  reservationRequest.CardMonth,
		CardYear:   reservationRequest.CardYear,
	}
}

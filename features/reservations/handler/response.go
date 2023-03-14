package handler

import (
	"groupproject3-airbnb-api/features/reservations"
	room "groupproject3-airbnb-api/features/rooms/handler"
	user "groupproject3-airbnb-api/features/user/handler"
	"reflect"
)

type ReservationResponse struct {
	Id         uint              `json:"id"`
	UserId     uint              `json:"user_id"`
	RoomId     uint              `json:"room_id"`
	DateStart  string            `json:"date_start"`
	DateEnd    string            `json:"date_end"`
	Duration   int               `json:"duration"`
	TotalPrice int               `json:"total_price"`
	CardName   string            `json:"card_name"`
	CardNumber string            `json:"card_number"`
	CardCvv    string            `json:"card_cvv"`
	CardMonth  string            `json:"card_month"`
	CardYear   string            `json:"card_yaer"`
	User       user.UserResponse `json:"user,omitempty"`
	Room       room.RoomResponse `json:"room,omitempty"`
}

func ReservationEntityToReservationResponse(reservationEntity reservations.ReservationEntity) ReservationResponse {
	result := ReservationResponse{
		Id:         reservationEntity.Id,
		UserId:     reservationEntity.UserId,
		RoomId:     reservationEntity.RoomId,
		DateStart:  reservationEntity.DateStart,
		DateEnd:    reservationEntity.DateEnd,
		Duration:   reservationEntity.Duration,
		TotalPrice: reservationEntity.TotalPrice,
		CardName:   reservationEntity.CardName,
		CardNumber: reservationEntity.CardNumber,
		CardCvv:    reservationEntity.CardCvv,
		CardMonth:  reservationEntity.CardMonth,
		CardYear:   reservationEntity.CardYear,
	}

	if !reflect.ValueOf(reservationEntity.User).IsZero() {
		result.User = user.UserResponse{
			ID:    reservationEntity.UserId,
			Name:  reservationEntity.User.Name,
			Email: reservationEntity.User.Email,
		}
	}

	if !reflect.ValueOf(reservationEntity.Room).IsZero() {
		result.Room = room.RoomResponse{
			RoomName: reservationEntity.Room.RoomName,
		}
	}

	return result
}

func ListReservationEntityToReservationResponse(reservationEntity []reservations.ReservationEntity) []ReservationResponse {
	var dataResponses []ReservationResponse
	for _, v := range reservationEntity {
		dataResponses = append(dataResponses, ReservationEntityToReservationResponse(v))
	}
	return dataResponses
}

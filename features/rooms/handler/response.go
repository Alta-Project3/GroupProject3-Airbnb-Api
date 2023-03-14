package handler

import (
	"groupproject3-airbnb-api/features/rooms"
	"groupproject3-airbnb-api/features/user/handler"
	"reflect"
)

type RoomResponse struct {
	Id          uint                 `json:"id"`
	UserId      uint                 `json:"user_id"`
	RoomName    string               `json:"room_name"`
	Price       int                  `json:"price"`
	Description string               `json:"description"`
	Latitude    float64              `json:"latitude"`
	Longitude   float64              `json:"longitude"`
	User        handler.UserResponse `json:"user,omitempty"`
}

func RoomEntityToRoomResponse(roomEntity rooms.RoomEntity) RoomResponse {
	result := RoomResponse{
		Id:          roomEntity.Id,
		UserId:      roomEntity.UserId,
		RoomName:    roomEntity.RoomName,
		Price:       roomEntity.Price,
		Description: roomEntity.Description,
		Latitude:    roomEntity.Latitude,
		Longitude:   roomEntity.Longitude,
	}

	if !reflect.ValueOf(roomEntity.User).IsZero() {
		result.User = handler.UserResponse{
			ID:    roomEntity.UserId,
			Name:  roomEntity.User.Name,
			Email: roomEntity.User.Email,
		}
	}

	return result
}

func ListRoomEntityToRoomResponse(roomEntity []rooms.RoomEntity) []RoomResponse {
	var dataResponses []RoomResponse
	for _, v := range roomEntity {
		dataResponses = append(dataResponses, RoomEntityToRoomResponse(v))
	}
	return dataResponses
}

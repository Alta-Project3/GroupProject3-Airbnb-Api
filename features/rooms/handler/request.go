package handler

import "groupproject3-airbnb-api/features/rooms"

type RoomRequest struct {
	UserId      uint    `json:"user_id" form:"user_id"`
	RoomName    string  `json:"room_name" form:"room_name"`
	Price       int     `json:"price" form:"price"`
	Description string  `json:"description" form:"description"`
	Latitude    float64 `json:"latitude" form:"latitude"`
	Longitude   float64 `json:"longitude" form:"longitude"`
}

func RoomRequestToRoomEntity(roomRequest *RoomRequest) rooms.RoomEntity {
	return rooms.RoomEntity{
		RoomName:    roomRequest.RoomName,
		Price:       roomRequest.Price,
		Description: roomRequest.Description,
		Latitude:    roomRequest.Latitude,
		Longitude:   roomRequest.Longitude,
	}
}

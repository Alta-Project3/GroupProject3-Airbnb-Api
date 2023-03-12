package data

import (
	"groupproject3-airbnb-api/features/rooms"
	"groupproject3-airbnb-api/features/user"
	"groupproject3-airbnb-api/features/user/data"
	"reflect"

	"gorm.io/gorm"
)

type Room struct {
	gorm.Model
	UserId      uint
	User        *data.User
	RoomName    string
	Price       int
	Description string
	Latitude    float64
	Longitude   float64
}

func RoomEntityToRoom(roomEntity rooms.RoomEntity) Room {
	result := Room{
		UserId:      roomEntity.UserId,
		RoomName:    roomEntity.RoomName,
		Price:       roomEntity.Price,
		Description: roomEntity.Description,
		Latitude:    roomEntity.Latitude,
		Longitude:   roomEntity.Longitude,
	}

	return result
}

func RoomToRoomEntity(room Room) rooms.RoomEntity {
	result := rooms.RoomEntity{
		Id:          room.ID,
		UserId:      room.UserId,
		RoomName:    room.RoomName,
		Price:       room.Price,
		Description: room.Description,
		Latitude:    room.Latitude,
		Longitude:   room.Longitude,
	}

	if !reflect.ValueOf(room.User).IsZero() {
		result.User = user.Core{
			Name:  room.User.Name,
			Email: room.User.Email,
		}
	}

	return result
}

func ListRoomToRoomEntity(room []Room) (roomEntity []rooms.RoomEntity) {
	for _, v := range room {
		roomEntity = append(roomEntity, RoomToRoomEntity(v))
	}
	return roomEntity
}

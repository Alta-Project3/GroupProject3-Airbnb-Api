package data

import (
	"errors"
	"groupproject3-airbnb-api/features/rooms"
	"groupproject3-airbnb-api/features/user/data"

	"gorm.io/gorm"
)

type query struct {
	db *gorm.DB
}

func New(db *gorm.DB) rooms.RoomDataInterface {
	return &query{
		db: db,
	}
}

func (q *query) SelectAll() ([]rooms.RoomEntity, error) {
	var rooms []Room
	if err := q.db.Preload("User").Find(&rooms); err.Error != nil {
		return nil, err.Error
	}
	return ListRoomToRoomEntity(rooms), nil
}

func (q *query) SelectById(id uint) (rooms.RoomEntity, error) {
	var room Room
	if err := q.db.Preload("User").First(&room, id); err.Error != nil {
		return rooms.RoomEntity{}, err.Error
	}
	return RoomToRoomEntity(room), nil
}

func (q *query) SelectByUserId(userId uint) ([]rooms.RoomEntity, error) {
	var room []Room
	err := q.db.Preload("User").
		Select("rooms.*").
		Where("user_id", userId).InnerJoins("User").
		Find(&room)

	if err.Error != nil {
		return nil, err.Error
	}
	return ListRoomToRoomEntity(room), nil
}

func (q *query) Store(roomEntity rooms.RoomEntity, userId uint) (uint, error) {
	var user data.User
	if err := q.db.First(&user, userId); err.Error != nil {
		return 0, err.Error
	}

	if user.Role == "User" {
		return 0, errors.New("only hosting role can create the room")
	}

	var room = RoomEntityToRoom(roomEntity)
	if err := q.db.Create(&room); err.Error != nil {
		return 0, err.Error
	}

	return room.ID, nil
}

func (q *query) Edit(roomEntity rooms.RoomEntity, id uint) error {
	var room = RoomEntityToRoom(roomEntity)
	if err := q.db.Where("id", id).Updates(&room); err.Error != nil {
		return err.Error
	}
	return nil
}

func (q *query) Destroy(id uint) error {
	var room Room
	if err := q.db.Delete(&room, id); err.Error != nil {
		return err.Error
	}

	return nil
}

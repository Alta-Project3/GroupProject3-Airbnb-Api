package data

import (
	"groupproject3-airbnb-api/features/rooms"

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

func (q *query) SelectByUserId(user_id uint) ([]rooms.RoomEntity, error) {
	var room []Room
	err := q.db.Preload("User").
		Select("rooms.*").
		Where("user_id", user_id).
		// InnerJoins("User").
		First(&room)

	if err.Error != nil {
		return nil, err.Error
	}
	return ListRoomToRoomEntity(room), nil
}

func (q *query) Store(roomEntity rooms.RoomEntity) (uint, error) {
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

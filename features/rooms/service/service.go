package service

import (
	"groupproject3-airbnb-api/features/rooms"

	"github.com/go-playground/validator/v10"
)

type roomService struct {
	Data     rooms.RoomDataInterface
	validate *validator.Validate
}

func New(data rooms.RoomDataInterface) rooms.RoomServiceInterface {
	return &roomService{
		Data:     data,
		validate: validator.New(),
	}
}

func (s *roomService) GetAll() ([]rooms.RoomEntity, error) {
	return s.Data.SelectAll()
}

func (s *roomService) GetById(id uint) (rooms.RoomEntity, error) {
	return s.Data.SelectById(id)
}

func (s *roomService) GetByUserId(user_id uint) ([]rooms.RoomEntity, error) {
	return s.Data.SelectByUserId(user_id)
}

func (s *roomService) Create(roomEntity rooms.RoomEntity) (rooms.RoomEntity, error) {
	s.validate = validator.New()
	errValidate := s.validate.StructExcept(roomEntity, "User")
	if errValidate != nil {
		return rooms.RoomEntity{}, errValidate
	}

	room_id, err := s.Data.Store(roomEntity)
	if err != nil {
		return rooms.RoomEntity{}, err
	}

	return s.Data.SelectById(room_id)
}

func (s *roomService) Update(roomEntity rooms.RoomEntity, id uint) (rooms.RoomEntity, error) {
	if checkDataExist, err := s.Data.SelectById(id); err != nil {
		return checkDataExist, err
	}

	err := s.Data.Edit(roomEntity, id)
	if err != nil {
		return rooms.RoomEntity{}, err
	}
	return s.Data.SelectById(id)
}

func (s *roomService) Delete(id uint) error {
	if _, err := s.Data.SelectById(id); err != nil {
		return err
	}

	return s.Data.Destroy(id)
}

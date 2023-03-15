package service

import (
	"errors"
	"groupproject3-airbnb-api/features/rooms"
	"groupproject3-airbnb-api/helper"
	"mime/multipart"

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

func (s *roomService) GetByUserId(userId, userIdLogin uint) ([]rooms.RoomEntity, error) {
	if userId != userIdLogin {
		return nil, errors.New("not allowed to access this user id")
	}
	return s.Data.SelectByUserId(userId)
}

func (s *roomService) Create(roomEntity rooms.RoomEntity, userId uint, fileData multipart.FileHeader) (rooms.RoomEntity, error) {
	//cek if user not host
	s.validate = validator.New()
	errValidate := s.validate.StructExcept(roomEntity, "User")
	if errValidate != nil {
		return rooms.RoomEntity{}, errValidate
	}

	roomEntity.UserId = userId
	url, err := helper.GetUrlImagesFromAWS(fileData)
	if err != nil {
		return rooms.RoomEntity{}, errors.New("validate: " + err.Error())
	}
	roomEntity.RoomPicture = url
	room_id, err := s.Data.Store(roomEntity, userId)
	if err != nil {
		return rooms.RoomEntity{}, err
	}

	return s.Data.SelectById(room_id)
}

func (s *roomService) Update(roomEntity rooms.RoomEntity, id, userId uint) (rooms.RoomEntity, error) {
	checkDataExist, err1 := s.Data.SelectById(id)
	if err1 != nil {
		return checkDataExist, err1
	}

	if checkDataExist.UserId != userId {
		return checkDataExist, errors.New("not allowed to access this Id")
	}

	err := s.Data.Edit(roomEntity, id)
	if err != nil {
		return rooms.RoomEntity{}, err
	}
	return s.Data.SelectById(id)
}

func (s *roomService) Delete(id, userId uint) error {
	checkDataExist, err := s.Data.SelectById(id)
	if err != nil {
		return err
	}

	if checkDataExist.UserId != userId {
		return errors.New("not allowed to access this Id")
	}

	return s.Data.Destroy(id)
}

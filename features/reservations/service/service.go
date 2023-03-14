package service

import (
	"errors"
	"groupproject3-airbnb-api/features/reservations"
	"groupproject3-airbnb-api/helper"

	"github.com/go-playground/validator/v10"
)

type ReservationService struct {
	Data     reservations.ReservationDataInterface
	validate *validator.Validate
}

func New(data reservations.ReservationDataInterface) reservations.ReservationServiceInterface {
	return &ReservationService{
		Data:     data,
		validate: validator.New(),
	}
}

func (s *ReservationService) CheckAvailability(reservationEntity reservations.ReservationEntity) (bool, error) {
	s.validate = validator.New()
	errValidate := s.validate.StructExcept(reservationEntity, "User", "Room")
	if errValidate != nil {
		return false, errValidate
	}

	dateStart, checkDateStart := helper.IsDate(reservationEntity.DateStart)
	if !checkDateStart {
		return false, errors.New("not valid date start. format date, ex : 2006-02-25")
	}

	dateEnd, checkDateEnd := helper.IsDate(reservationEntity.DateEnd)
	if !checkDateEnd {
		return false, errors.New("not valid date end. format date, ex : 2006-02-25")
	}

	if helper.FormatDate(dateEnd).Before(helper.FormatDate(dateStart)) {
		return false, errors.New("error range of date, date end must be after date start")
	}

	data, err := s.Data.SelectyRoomAndDateRange(reservationEntity)
	if err != nil {
		return false, err
	}

	if len(data) > 0 {
		return false, errors.New("date not available")
	}

	return true, nil
}

// Create implements reservations.ReservationServiceInterface
func (*ReservationService) Create(reservationEntity reservations.ReservationEntity) (reservations.ReservationEntity, error) {
	// dateStart, checkDateStart := helper.IsDate(reservationEntity.DateStart)
	// if !checkDateStart {
	// 	return false, errors.New("not valid date start format date, ex : 2006-02-25")
	// }

	// dateEnd, checkDateEnd := helper.IsDate(reservationEntity.DateEnd)
	// if !checkDateEnd {
	// 	return false, errors.New("not valid date graduate format date, ex : 2006-02-25")
	// }

	// if helper.FormatDate(dateEnd).Before(helper.FormatDate(dateStart)) {
	// 	return false, errors.New("error range of date, date end must be after date start")
	// }
	panic("unimplemented")
}

// GetByRoomAndDateRange implements reservations.ReservationServiceInterface
func (s *ReservationService) GetReservation(userId uint) ([]reservations.ReservationEntity, error) {
	return s.Data.SelectyReservation(uint(userId))
}

package service

import (
	"errors"
	"groupproject3-airbnb-api/features/reservations"
	"groupproject3-airbnb-api/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// func TestCreate(t *testing.T) {
// 	repo := mocks.NewReservationDataInterface(t)
// 	inputData := reservations.ReservationEntity{Id: 1, RoomId: 1, DateStart: "2023-03-16", DateEnd: "2023-03-19", TotalPrice: 2500000}
// 	ResData := reservations.ReservationEntity{Id: 1, RoomId: 1, DateStart: "2023-03-16", DateEnd: "2023-03-19", TotalPrice: 2500000}

// 	t.Run("success make reservation", func(t *testing.T) {
// 		repo.On("Store", inputData,"SelectById", uint(1)).Return(uint(1), nil).Once()
// 		repo.On("SelectById", uint(1)).Return(ResData, nil)
// 		srv := New(repo)

// 	})

// }
func TestCreate(t *testing.T) {
	repo := mocks.NewReservationDataInterface(t)

	inputData := reservations.ReservationEntity{Id: 1, RoomId: 1, DateStart: "2023-03-16", DateEnd: "2023-03-19", TotalPrice: 2500000}
	srv := New(repo)

	// resData := reservations.ReservationEntity{Id: 1, RoomId: 1, DateStart: "2023-03-16", DateEnd: "2023-03-19", TotalPrice: 2500000}

	t.Run("success make reservation", func(t *testing.T) {
		repo.On("Store", mock.Anything).Return(uint(1), nil)
		res, err := srv.Create(inputData)

		assert.Nil(t, err)
		assert.NotEqual(t, inputData, res)
		repo.AssertExpectations(t)
	})

	t.Run("not valid date start frmat", func(t *testing.T) {
		repo.On("Store", mock.Anything).Return(uint(1), errors.New("not valid date start format"))
		res, err := srv.Create(inputData)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "valid")
		assert.Equal(t, reservations.ReservationEntity{}, res)
		repo.AssertExpectations(t)
	})

}

func TestGetReservation(t *testing.T) {
	repo := mocks.NewReservationDataInterface(t)

	resData := []reservations.ReservationEntity{{Id: 1, RoomId: 1, DateStart: "2023-03-16", DateEnd: "2023-03-19", TotalPrice: 2500000}}
	srv := New(repo)

	t.Run("success get reservation", func(t *testing.T) {
		repo.On("SelectyReservation", mock.Anything).Return(resData, nil)
		res, err := srv.GetReservation(uint(1))

		assert.Nil(t, err)
		assert.Equal(t, len(resData), len(res))
		repo.AssertExpectations(t)
	})

}

func TestGetById(t *testing.T) {
	repo := mocks.NewReservationDataInterface(t)

	resData := []reservations.ReservationEntity{{Id: 1, RoomId: 1, DateStart: "2023-03-16", DateEnd: "2023-03-19", TotalPrice: 2500000}}
	srv := New(repo)

	t.Run("success get reservation", func(t *testing.T) {
		repo.On("SelectByRoomId", mock.Anything).Return(resData, nil)
		res, err := srv.GetByRoomId(uint(1))

		assert.Nil(t, err)
		assert.Equal(t, len(resData), len(res))
		repo.AssertExpectations(t)
	})

}

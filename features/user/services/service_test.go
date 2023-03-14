package services

import (
	"errors"
	"groupproject3-airbnb-api/features/user"
	"groupproject3-airbnb-api/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestRegister(t *testing.T) {
	repo := mocks.NewUserData(t)
	inputData := user.Core{Name: "Alif", Email: "alif@gmail.com", Phone: "081234"}
	t.Run("Success register", func(t *testing.T) {
		repo.On("Register", mock.Anything).Return(nil).Once()
		srv := New(repo)
		err := srv.Register(inputData)
		assert.Nil(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("invalid validation", func(t *testing.T) {
		wrongInput := user.Core{Name: "Alif%&*", Password: "123a#$%"}
		// repo.On("Register", uint(1), mock.Anything).Return(user.Core{}, errors.New("email duplicated")).Once()
		srv := New(repo)
		err := srv.Register(wrongInput)
		assert.ErrorContains(t, err, "validate")
		repo.AssertExpectations(t)

	})

	t.Run("Duplicated", func(t *testing.T) {
		repo.On("Register", mock.Anything).Return(errors.New("duplicated")).Once()
		srv := New(repo)
		err := srv.Register(inputData)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "email already registered")
		repo.AssertExpectations(t)
	})

	t.Run("internal server error", func(t *testing.T) {
		repo.On("Register", mock.Anything).Return(errors.New("There is a problem with the server")).Once()
		srv := New(repo)
		err := srv.Register(inputData)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "server")
		repo.AssertExpectations(t)
	})
}

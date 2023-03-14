package services

import (
	"errors"
	"groupproject3-airbnb-api/features/user"
	"groupproject3-airbnb-api/helper"
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

func TestLogin(t *testing.T) {
	repo := mocks.NewUserData(t)
	inputEmail := "alif@gmail.com"
	passwordHashed := helper.GeneratePassword("123")
	resData := user.Core{ID: uint(1), Name: "Alif", Email: "alif@gmail.com", Password: passwordHashed}

	t.Run("login success", func(t *testing.T) {
		repo.On("Login", inputEmail).Return(resData, nil).Once()
		srv := New(repo)
		token, res, err := srv.Login(inputEmail, "123")
		assert.Nil(t, err)
		assert.NotEmpty(t, token)
		assert.Equal(t, resData.Name, res.Name)
		repo.AssertExpectations(t)
	})

	t.Run("account not found", func(t *testing.T) {
		repo.On("Login", inputEmail).Return(user.Core{}, errors.New("data not found")).Once()
		srv := New(repo)
		token, res, err := srv.Login(inputEmail, "123")
		assert.NotNil(t, token)
		assert.ErrorContains(t, err, "not")
		assert.Empty(t, token)
		assert.Equal(t, uint(0), res.ID)
		repo.AssertExpectations(t)
	})

	t.Run("password not matched", func(t *testing.T) {
		inputEmail := "lip23"
		repo.On("Login", inputEmail).Return(resData, nil).Once()
		srv := New(repo)
		_, res, err := srv.Login(inputEmail, "342")
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "password")
		assert.Empty(t, nil)
		assert.Equal(t, uint(0), res.ID)
		repo.AssertExpectations(t)
	})

	t.Run("server problem", func(t *testing.T) {
		inputEmail := "alif23"
		repo.On("Login", inputEmail).Return(user.Core{}, errors.New("There is a problem with the server")).Once()

		srv := New(repo)
		token, res, err := srv.Login(inputEmail, "342")
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "server")
		assert.Empty(t, token)
		assert.Equal(t, user.Core{}, res)
		repo.AssertExpectations(t)
	})
}

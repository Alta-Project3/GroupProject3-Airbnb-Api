package services

import (
	"errors"
	"groupproject3-airbnb-api/features/feedback"
	"groupproject3-airbnb-api/helper"
	"groupproject3-airbnb-api/mocks"
	"testing"

	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	repo := mocks.NewFeedbackData(t)

	inputData := feedback.Core{ID: 1, RoomID: 1, Rating: 3, Feedback: "nyaman"}
	resData := feedback.Core{ID: 1, RoomID: 1, Rating: 3, Feedback: "nyaman"}
	t.Run("success add feedback", func(t *testing.T) {
		repo.On("Create", uint(1), uint(1), inputData).Return(resData, nil).Once()
		srv := New(repo)
		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.Create(pToken, uint(1), inputData)
		assert.Nil(t, err)
		assert.Equal(t, resData.ID, res.ID)
		repo.AssertExpectations(t)
	})

	t.Run("user not found", func(t *testing.T) {
		repo.On("Create", uint(1), uint(1), inputData).Return(feedback.Core{}, errors.New("user not found")).Once()
		srv := New(repo)
		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.Create(pToken, uint(1), inputData)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "server")
		assert.Equal(t, feedback.Core{}, res)
		repo.AssertExpectations(t)
	})

	t.Run("data not found", func(t *testing.T) {
		repo.On("Create", uint(1), uint(1), inputData).Return(feedback.Core{}, errors.New("data not found")).Once()
		srv := New(repo)
		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.Create(pToken, uint(1), inputData)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "not found")
		assert.Equal(t, feedback.Core{}, res)
		repo.AssertExpectations(t)
	})

	t.Run("server problem", func(t *testing.T) {
		repo.On("Create", uint(1), uint(1), inputData).Return(feedback.Core{}, errors.New("server problem")).Once()
		srv := New(repo)
		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.Create(pToken, uint(1), inputData)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "server")
		assert.Equal(t, feedback.Core{}, res)
		repo.AssertExpectations(t)
	})

}

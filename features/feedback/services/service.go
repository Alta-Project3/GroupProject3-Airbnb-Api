package services

import (
	"errors"
	"groupproject3-airbnb-api/features/feedback"
	"groupproject3-airbnb-api/helper"
	"log"
)

type feedbackUseCase struct {
	qry feedback.FeedbackData
}

func New(fd feedback.FeedbackData) feedback.FeedbackService {
	return &feedbackUseCase{
		qry: fd,
	}
}

func (fuc *feedbackUseCase) Create(token interface{}, newFeedback feedback.Core) (feedback.Core, error) {
	userID := helper.ExtractToken(token)

	if userID <= 0 {
		return feedback.Core{}, errors.New("user not found")
	}
	res, err := fuc.qry.Create(uint(userID), newFeedback)

	if err != nil {
		log.Println("cannot post book", err.Error())
		return feedback.Core{}, errors.New("server error")
	}

	return res, nil
}

// GetAll implements feedback.FeedbackService
func (*feedbackUseCase) GetAll() ([]feedback.Core, error) {
	panic("unimplemented")
}

// GetByID implements feedback.FeedbackService
func (*feedbackUseCase) GetByID(token interface{}, feedbackID uint) (feedback.Core, error) {
	panic("unimplemented")
}

// Update implements feedback.FeedbackService
func (*feedbackUseCase) Update(token interface{}, feedbackID uint) (feedback.Core, error) {
	panic("unimplemented")
}

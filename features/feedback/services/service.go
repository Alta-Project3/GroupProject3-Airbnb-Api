package services

import (
	"errors"
	"groupproject3-airbnb-api/features/feedback"
	"groupproject3-airbnb-api/helper"
	"log"
	"strings"
)

type feedbackUseCase struct {
	qry feedback.FeedbackData
}

func New(fd feedback.FeedbackData) feedback.FeedbackService {
	return &feedbackUseCase{
		qry: fd,
	}
}

func (fuc *feedbackUseCase) Create(token interface{}, roomID uint, newFeedback feedback.Core) (feedback.Core, error) {
	userID := helper.ExtractToken(token)

	if userID <= 0 {
		return feedback.Core{}, errors.New("user not found")
	}
	res, err := fuc.qry.Create(uint(userID), roomID, newFeedback)

	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "data not found"
		} else {
			msg = "server problem"
		}
		log.Println("error add query in service: ", err.Error())
		return feedback.Core{}, errors.New(msg)
	}
	return res, nil
}

func (fuc *feedbackUseCase) GetUserFeedback(token interface{}) ([]feedback.Core, error) {
	userID := helper.ExtractToken(token)
	res, err := fuc.qry.GetUserFeedback(uint(userID))
	if err != nil {
		log.Println("query error", err.Error())
		return []feedback.Core{}, errors.New("query error, problem with server")
	}
	return res, nil
}

func (fuc *feedbackUseCase) GetByID(token interface{}, feedbackID uint) (feedback.Core, error) {
	userID := helper.ExtractToken(token)
	res, err := fuc.qry.GetByID(uint(userID), feedbackID)
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "feedback not found"
		} else {
			msg = "there is a problem with the server"
		}
		return feedback.Core{}, errors.New(msg)
	}
	return res, nil
}

// Update implements feedback.FeedbackService
func (*feedbackUseCase) Update(token interface{}, feedbackID uint) (feedback.Core, error) {
	panic("unimplemented")
}

package data

import (
	"groupproject3-airbnb-api/features/feedback"

	"gorm.io/gorm"
)

type feedbackQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) feedback.FeedbackData {
	return &feedbackQuery{
		db: db,
	}
}

func (fq *feedbackQuery) Create(userID uint, newFeedback feedback.Core) (feedback.Core, error) {
	cnv := CoreToData(newFeedback)
	cnv.UserID = uint(userID)
	err := fq.db.Create(&cnv).Error

	if err != nil {
		return feedback.Core{}, err
	}

	newFeedback.ID = cnv.ID
	newFeedback.UserID = cnv.UserID

	return newFeedback, nil
}

// GetAll implements feedback.FeedbackData
func (*feedbackQuery) GetAll() ([]feedback.Core, error) {
	panic("unimplemented")
}

// GetByID implements feedback.FeedbackData
func (*feedbackQuery) GetByID(userID uint, feedbackID uint) (feedback.Core, error) {
	panic("unimplemented")
}

// Update implements feedback.FeedbackData
func (*feedbackQuery) Update(userID uint, feedBackID uint) (feedback.Core, error) {
	panic("unimplemented")
}

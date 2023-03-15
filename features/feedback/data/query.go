package data

import (
	"errors"
	"groupproject3-airbnb-api/features/feedback"
	"log"

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

func (fq *feedbackQuery) Create(userID uint, roomID uint, newFeedback feedback.Core) (feedback.Core, error) {
	room := Room{}

	err := fq.db.Where("id=?", roomID).First(&room).Error
	if err != nil {
		log.Println("query error", err.Error())
		return feedback.Core{}, errors.New("server error")
	}

	cnv := CoreToData(newFeedback)
	cnv.RoomID = room.ID
	cnv.UserID = userID
	err = fq.db.Create(&cnv).Error
	if err != nil {
		log.Println("query error", err.Error())
		return feedback.Core{}, errors.New("server error")
	}
	result := DataToCore(cnv)
	return result, nil
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

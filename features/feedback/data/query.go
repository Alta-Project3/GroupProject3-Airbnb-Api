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

func (fq *feedbackQuery) GetUserFeedback(userID uint) ([]feedback.Core, error) {
	res := []Feedback{}
	err := fq.db.Where("user_id = ?", userID).Find(&res).Error

	if err != nil {
		log.Println("query error", err.Error())
		return []feedback.Core{}, errors.New("server error")
	}

	result := []feedback.Core{}
	for i := 0; i < len(res); i++ {
		result = append(result, DataToCore(res[i]))
		// cari data user berdasarkan cart user_id

		room := Room{}
		err = fq.db.Where("id = ?", res[i].RoomID).First(&room).Error
		if err != nil {
			log.Println("query error", err.Error())
			return []feedback.Core{}, errors.New("server error")
		}
		user := User{}
		err = fq.db.Where("id = ?", room.UserID).First(&user).Error
		if err != nil {
			log.Println("query error", err.Error())
			return []feedback.Core{}, errors.New("server error")
		}
		result[i].ID = room.ID
		result[i].Rating = room.Rating
		result[i].Feedback = room.Feedback
	}
	return result, nil
}

// GetByID implements feedback.FeedbackData
func (*feedbackQuery) GetByID(userID uint, feedbackID uint) (feedback.Core, error) {
	panic("unimplemented")
}

// Update implements feedback.FeedbackData
func (*feedbackQuery) Update(userID uint, feedBackID uint) (feedback.Core, error) {
	panic("unimplemented")
}

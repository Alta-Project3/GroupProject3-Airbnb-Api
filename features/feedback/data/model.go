package data

import (
	"groupproject3-airbnb-api/features/feedback"
	user "groupproject3-airbnb-api/features/user/data"

	"gorm.io/gorm"
)

type Feedback struct {
	gorm.Model
	UserID        uint
	User          user.User
	ReservationID uint
	RoomID        uint
	Rating        float64
	Feedback      string
}

func DataToCore(data Feedback) feedback.Core {
	return feedback.Core{
		ID:            data.ID,
		UserID:        data.UserID,
		ReservationID: data.ReservationID,
		RoomID:        data.RoomID,
		Rating:        data.Rating,
		Feedback:      data.Feedback,
	}
}

func CoreToData(data feedback.Core) Feedback {
	return Feedback{
		Model:         gorm.Model{ID: data.ID},
		UserID:        data.UserID,
		ReservationID: data.ReservationID,
		RoomID:        data.RoomID,
		Rating:        data.Rating,
		Feedback:      data.Feedback,
	}
}

package feedback

import (
	"groupproject3-airbnb-api/features/user"
	"time"
)

type FeedbackEntity struct {
	Id            uint
	UserId        uint `validate:"required"`
	User          user.Core
	ReservationId uint    `validate:"required"`
	RoomId        uint    `validate:"required"`
	Rating        float64 `validate:"required"`
	Feedback      string  `validate:"required"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type FeedbackServiceInterface interface {
	GetAll() ([]FeedbackEntity, error)
	GetById(id uint) (FeedbackEntity, error)
	GetByRoomId(roomId uint) ([]FeedbackEntity, error)
	Create(feedbackEntity FeedbackEntity) (FeedbackEntity, error)
	Update(feedbackEntity FeedbackEntity, id uint) (FeedbackEntity, error)
	Delete(id uint) error
}

type FeedbackDataInterface interface {
	SelectAll() ([]FeedbackEntity, error)
	SelectById(id uint) (FeedbackEntity, error)
	SelectByRoomId(id uint) (FeedbackEntity, error)
	Store(feedbackEntity FeedbackEntity) (uint, error)
	Edit(feedbackEntity FeedbackEntity, id uint) error
	Destroy(id uint) error
}

package feedback

import (
	"groupproject3-airbnb-api/features/user"

	"github.com/labstack/echo/v4"
)

type Core struct {
	ID            uint
	UserId        uint `validate:"required"`
	User          user.Core
	ReservationID uint    `validate:"required"`
	RoomID        uint    `validate:"required"`
	Rating        float64 `validate:"required"`
	Feedback      string  `validate:"required"`
}

type FeedbackHandler interface {
	Create() echo.HandlerFunc
	GetAll() echo.HandlerFunc
	GetByID() echo.HandlerFunc
	Update() echo.HandlerFunc
}

type FeedbackService interface {
	Create(token interface{}, newFeedback Core) (Core, error)
	GetAll() ([]Core, error)
	GetByID(token interface{}, feedbackID uint) (Core, error)
	Update(token interface{}, feedbackID uint) (Core, error)
}

type FeedbackData interface {
	Create(userID uint, newFeedback Core) (Core, error)
	GetAll() ([]Core, error)
	GetByID(userID uint, feedbackID uint) (Core, error)
	Update(userID uint, feedBackID uint) (Core, error)
}

package feedback

import (
	"github.com/labstack/echo/v4"
)

type Core struct {
	ID            uint
	UserID        uint `validate:"required"`
	User          User
	ReservationID uint    `validate:"required"`
	RoomID        uint    `validate:"required"`
	Rating        float64 `validate:"required"`
	Feedback      string  `validate:"required"`
}

type User struct {
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
}

type FeedbackHandler interface {
	Create() echo.HandlerFunc
	GetAll() echo.HandlerFunc
	GetByID() echo.HandlerFunc
	Update() echo.HandlerFunc
}

type FeedbackService interface {
	Create(token interface{}, roomID uint, newFeedback Core) (Core, error)
	GetUserFeedback(token interface{}) ([]Core, error)
	GetByID(token interface{}, feedbackID uint) (Core, error)
	Update(token interface{}, feedbackID uint) (Core, error)
}

type FeedbackData interface {
	Create(userID uint, roomID uint, newFeedback Core) (Core, error)
	GetUserFeedback(userID uint) ([]Core, error)
	GetByID(userID uint, feedbackID uint) (Core, error)
	Update(userID uint, feedBackID uint) (Core, error)
}

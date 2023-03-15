package handler

import (
	"groupproject3-airbnb-api/features/feedback"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type feedbackControll struct {
	srv feedback.FeedbackService
}

func New(srv feedback.FeedbackService) feedback.FeedbackHandler {
	return &feedbackControll{
		srv: srv,
	}
}

func (fc *feedbackControll) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Get("user")
		input := CreateFeedbackRequest{}

		err := c.Bind(&input)
		if err != nil {
			return c.JSON(http.StatusBadRequest, "input format incorrect")
		}

		res, err := fc.srv.Create(token, input.RoomID, *ReqToCore(input))
		if err != nil {
			log.Println("error running create feedback service: ", err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": "server problem"})
		}
		return c.JSON(http.StatusCreated, map[string]interface{}{
			"data":    ToFeedbackResponse(res),
			"message": "success add feedback",
		})

	}
}

// GetAll implements feedback.FeedbackHandler
func (*feedbackControll) GetAll() echo.HandlerFunc {
	panic("unimplemented")
}

// GetByID implements feedback.FeedbackHandler
func (*feedbackControll) GetByID() echo.HandlerFunc {
	panic("unimplemented")
}

// Update implements feedback.FeedbackHandler
func (*feedbackControll) Update() echo.HandlerFunc {
	panic("unimplemented")
}

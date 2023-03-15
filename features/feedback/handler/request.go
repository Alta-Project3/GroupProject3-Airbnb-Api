package handler

import "groupproject3-airbnb-api/features/feedback"

type CreateFeedbackRequest struct {
	RoomID   uint    `json:"room_id" form:"room_id"`
	Rating   float64 `json:"rating" form:"rating"`
	Feedback string  `json:"feedback" form:"feedback"`
}

func ReqToCore(data interface{}) *feedback.Core {
	res := feedback.Core{}

	switch data.(type) {
	case CreateFeedbackRequest:
		cnv := data.(CreateFeedbackRequest)
		res.Rating = cnv.Rating
		res.Feedback = cnv.Feedback
	default:
		return nil
	}

	return &res
}

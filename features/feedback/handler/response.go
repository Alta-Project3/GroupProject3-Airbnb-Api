package handler

import "groupproject3-airbnb-api/features/feedback"

type FeedbackResponse struct {
	ID       uint    `json:"id,omitempty"`
	RoomID   uint    `json:"room_id,omitempty"`
	Rating   float64 `json:"rating,omitempty"`
	Feedback string  `json:"feedback,omitempty"`
}

func ToFeedbackResponse(data feedback.Core) FeedbackResponse {
	return FeedbackResponse{
		ID:       data.ID,
		RoomID:   data.RoomID,
		Rating:   data.Rating,
		Feedback: data.Feedback,
	}
}

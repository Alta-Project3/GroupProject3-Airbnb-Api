package handler

import "groupproject3-airbnb-api/features/user"

type UserReponse struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func ToResponse(data user.Core) UserReponse {
	return UserReponse{
		ID:    data.ID,
		Name:  data.Name,
		Email: data.Email,
	}
}

type ProfileResponse struct {
	ID             uint   `json:"id"`
	ProfilePicture string `json:"profile_picture"`
	Name           string `json:"name"`
	Email          string `json:"email"`
	Phone          string `json:"phone"`
	Address        string `json:"address"`
	Role           string `json:"role"`
}

func ToProfileResponse(data user.Core) ProfileResponse {
	return ProfileResponse{
		ID:             data.ID,
		ProfilePicture: data.ProfilePicture,
		Name:           data.Name,
		Email:          data.Email,
		Phone:          data.Phone,
		Address:        data.Address,
		Role:           data.Role,
	}
}

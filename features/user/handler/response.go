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

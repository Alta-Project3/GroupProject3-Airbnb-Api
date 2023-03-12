package services

import (
	"errors"
	"groupproject3-airbnb-api/features/user"
	"groupproject3-airbnb-api/helper"
	"strings"
)

type userUseCase struct {
	qry user.UserData
}

func New(ud user.UserData) user.UserService {
	return &userUseCase{
		qry: ud,
	}
}

func (uuc *userUseCase) Register(newUser user.Core) (user.Core, error) {
	if len(newUser.Password) != 0 {
		//validation
		err := helper.RegistrationValidate(newUser)
		if err != nil {
			return user.Core{}, errors.New("validate: " + err.Error())
		}
	}
	hashed := helper.GeneratePassword(newUser.Password)
	newUser.Password = string(hashed)

	res, err := uuc.qry.Register(newUser)
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "duplicated") {
			msg = "email already registered"
		} else {
			msg = "server error"
		}
		return user.Core{}, errors.New(msg)
	}

	return res, nil
}

// Login implements user.UserService
func (*userUseCase) Login(email string, password string) (string, user.Core, error) {
	panic("unimplemented")
}

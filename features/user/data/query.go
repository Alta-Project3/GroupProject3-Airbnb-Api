package data

import (
	"errors"
	"groupproject3-airbnb-api/features/user"
	"log"

	"gorm.io/gorm"
)

type userQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) user.UserData {
	return &userQuery{
		db: db,
	}
}

func (uq *userQuery) Register(newUser user.Core) (user.Core, error) {
	dupEmail := CoreToData(newUser)
	err := uq.db.Where("email = ?", newUser.Email).First(&dupEmail).Error
	if err == nil {
		log.Println("duplicated")
		return user.Core{}, errors.New("email duplicated")
	}

	newUser.ProfilePicture = "https://cdn.pixabay.com/photo/2015/10/05/22/37/blank-profile-picture-973460_1280.png"
	newUser.Role = "User"

	cnv := CoreToData(newUser)
	err = uq.db.Create(&cnv).Error
	if err != nil {
		log.Println("query error", err.Error())
		return user.Core{}, errors.New("server error")
	}

	newUser.ID = cnv.ID
	return newUser, nil
}

func (uq *userQuery) Login(email string) (user.Core, error) {
	if email == "" {
		log.Println("data empty, query error")
		return user.Core{}, errors.New("email not allowed empty")
	}
	res := User{}
	if err := uq.db.Where("email = ?", email).First(&res).Error; err != nil {
		log.Println("login query error", err.Error())
		return user.Core{}, errors.New("data not found")
	}

	return ToCore(res), nil
}

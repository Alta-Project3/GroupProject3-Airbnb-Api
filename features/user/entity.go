package user

import (
	"mime/multipart"

	"github.com/labstack/echo/v4"
)

type Core struct {
	ID             uint
	ProfilePicture string
	Name           string
	Email          string
	Password       string
	Phone          string
	Address        string
	Role           string
}

type UserHandler interface {
	Register() echo.HandlerFunc
	Login() echo.HandlerFunc
	Profile() echo.HandlerFunc
	Update() echo.HandlerFunc
	Deactivate() echo.HandlerFunc
}

type UserService interface {
	Register(newUser Core) error
	Login(email, password string) (string, Core, error)
	Profile(token interface{}) (Core, error)
	Update(token interface{}, fileData multipart.FileHeader, updateData Core) (Core, error)
	Deactivate(token interface{}) error
}

type UserData interface {
	Register(newUser Core) error
	Login(email string) (Core, error)
	Profile(userID uint) (Core, error)
	Update(userID uint, updateData Core) (Core, error)
	Deactivate(userID uint) error
}

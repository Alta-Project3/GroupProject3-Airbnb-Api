package user

import "github.com/labstack/echo/v4"

type Core struct {
	ID             uint
	ProfilePicture string
	Name           string
	Email          string
	Password       string
	Phone          string
	Address        string
}

type UserHandler interface {
	Register() echo.HandlerFunc
	Login() echo.HandlerFunc
}

type UserService interface {
	Register(newUser Core) (Core, error)
	Login(email, password string) (string, Core, error)
}

type UserData interface {
	Register(newUser Core) (Core, error)
	Login(email string) (Core, error)
}

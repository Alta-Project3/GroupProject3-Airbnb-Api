package router

import (
	"groupproject3-airbnb-api/app/config"
	usrData "groupproject3-airbnb-api/features/user/data"
	usrHdl "groupproject3-airbnb-api/features/user/handler"
	usrSrv "groupproject3-airbnb-api/features/user/services"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func userRouter(db *gorm.DB, e *echo.Echo) {
	uData := usrData.New(db)
	uService := usrSrv.New(uData)
	uHandler := usrHdl.New(uService)

	// AUTH
	e.POST("/register", uHandler.Register())
	e.POST("/login", uHandler.Login())

	g := e.Group("/users")
	g.Use(echojwt.JWT([]byte(config.JWTKey)))
	g.GET("", uHandler.Profile())
	g.PUT("", uHandler.Update())
	g.DELETE("", uHandler.Deactivate())
}

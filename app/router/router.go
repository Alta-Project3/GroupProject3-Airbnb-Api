package router

import (
	usrData "groupproject3-airbnb-api/features/user/data"
	usrHdl "groupproject3-airbnb-api/features/user/handler"
	usrSrv "groupproject3-airbnb-api/features/user/services"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRouter(db *gorm.DB, e *echo.Echo) {
	uData := usrData.New(db)
	uService := usrSrv.New(uData)
	uHandler := usrHdl.New(uService)

	// AUTH
	e.POST("/register", uHandler.Register())
	e.POST("/login", uHandler.Login())
}

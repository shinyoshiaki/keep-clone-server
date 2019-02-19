package main

import (
	"keep-server/handler/memo/addmemo"
	"keep-server/handler/memo/getmemo"
	"keep-server/handler/user/login"
	"keep-server/handler/user/signup"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	e.POST("/user/signup", signup.SignUp)
	e.POST("/user/login", login.Login)
	e.POST("/memo/post", addmemo.Post)
	e.POST("memo/get", getmemo.Get)

	e.Start(":1323")
}

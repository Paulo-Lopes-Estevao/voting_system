package server

import (
	"gomq/user/route"

	"github.com/labstack/echo/v4"
)

func UserServer() {
	e := echo.New()

	route.UserRoute(e)

	e.Logger.Fatal(e.Start(":8001"))
}

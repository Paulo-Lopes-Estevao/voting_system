package main

import (
	"gomq/user/route"

	"github.com/labstack/echo/v4"
)

func main() {

	e := echo.New()

	route.UserRoute(e)

	e.Logger.Fatal(e.Start(":8000"))

}

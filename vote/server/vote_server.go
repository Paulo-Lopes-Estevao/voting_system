package server

import (
	"gomq/vote/route"

	"github.com/labstack/echo/v4"
)

func VoteServer() {
	e := echo.New()

	route.VoteRoute(e)

	e.Logger.Fatal(e.Start(":8002"))
}

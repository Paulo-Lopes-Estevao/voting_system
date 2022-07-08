package main

import (
	user_roouting "gomq/user/route"
	vote_routing "gomq/vote/route"

	"github.com/labstack/echo/v4"
)

func main() {

	e := echo.New()

	user_roouting.UserRoute(e)
	vote_routing.VoteRoute(e)

	e.Logger.Fatal(e.Start(":8000"))

}

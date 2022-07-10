package server

import (
	"gomq/vote/controller"
	"gomq/vote/factory"
	"gomq/vote/infra/db"
	"gomq/vote/route"

	"github.com/labstack/echo/v4"
)

func VoteServer() {
	e := echo.New()

	conn := db.ConnectDB()
	voteUsecaseFactory := factory.VoteUsecaseFactory(conn)
	voteController := controller.New(&voteUsecaseFactory)
	route.VoteRoute(e, voteController)

	e.Logger.Fatal(e.Start(":8002"))
}

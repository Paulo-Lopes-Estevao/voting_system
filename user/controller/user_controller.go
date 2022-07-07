package controller

import (
	"encoding/json"
	"fmt"
	"gomq/user/queue"
	"gomq/user/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserVote struct {
	Name string
	Vote string
}

func HomePageVoteController(c echo.Context) error {
	return c.Render(http.StatusOK, "user_vote.html", map[string]interface{}{})
}

func UserVoteController(c echo.Context) error {
	var user_vote UserVote
	user_vote.Name = c.FormValue("name")
	user_vote.Vote = c.FormValue("socialnetwork")

	_, err := usecase.UserUsecase(user_vote.Name)
	if err != nil {
		return c.Render(http.StatusOK, "user_vote.html", map[string]interface{}{
			"error": err.Error(),
		})
	}

	data, _ := json.Marshal(user_vote)
	fmt.Println(string(data))

	connection := queue.Connect()
	queue.UserMarkVote(data, "vote_ex", "vote_rk", connection)

	return c.Render(http.StatusOK, "user_vote.html", map[string]interface{}{})
}

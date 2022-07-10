package controller

import (
	"fmt"
	"gomq/vote/dto"
	"gomq/vote/queue"
	"gomq/vote/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

type IVoteController interface {
	ShowTheVotesController(c echo.Context) error
}

type voteController struct {
	Ivoteusecase usecase.IVoteUseCase
}

func New(Ivoteusecase usecase.IVoteUseCase) IVoteController {
	return &voteController{
		Ivoteusecase: Ivoteusecase,
	}
}

func (vc *voteController) ShowTheVotesController(c echo.Context) error {
	in := make(chan []byte)

	connection := queue.Connect()
	go queue.ShowVotes("vote_ex", "vote_rk", "direct", connection, in)

	var v dto.Vote
	go func() {
		for payload := range in {
			v.ParseJson(payload)
			fmt.Println(v)
			vc.Ivoteusecase.RegisterVotes(v.Name, v.Vote)
		}
	}()
	votes_list, _ := vc.Ivoteusecase.ShowVotes()
	return c.Render(http.StatusOK, "votes.html", map[string]interface{}{
		"data": votes_list,
	})
}

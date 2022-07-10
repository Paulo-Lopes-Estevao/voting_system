package repository

import "gomq/vote/entities"

type IVoteRepository interface {
	AddVotes(votes *entities.Vote) error
	ShowVotes() (*entities.Vote, error)
}

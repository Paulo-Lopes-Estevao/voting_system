package repository

import "gomq/vote/entities"

type IVoteRepository interface {
	AddVotes(name, vote string) error
	ShowVotes() entities.Vote
}

package usecase

import (
	"gomq/vote/entities"
	"gomq/vote/entities/repository"
)

type IVoteUseCase interface {
	RegisterVotes(name, vote string) (*entities.Vote, error)
	ShowVotes() ([]*entities.Vote, error)
}

type VoteUseCase struct {
	Irespositoy repository.IVoteRepository
}

func New(Irespo repository.IVoteRepository) IVoteUseCase {
	return &VoteUseCase{
		Irespositoy: Irespo,
	}
}

func (v *VoteUseCase) RegisterVotes(name, vote string) (*entities.Vote, error) {
	votes := &entities.Vote{
		Name: name,
		Vote: vote,
	}
	err := v.Irespositoy.AddVotes(votes)
	if err != nil {
		return nil, err
	}
	return votes, nil
}

func (v *VoteUseCase) ShowVotes() ([]*entities.Vote, error) {
	votes, err := v.Irespositoy.ShowVotes()
	if err != nil {
		return nil, err
	}
	return votes, nil
}

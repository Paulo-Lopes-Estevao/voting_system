package repository

import (
	"gomq/vote/entities"

	"github.com/jinzhu/gorm"
)

type VoteRepository struct {
	Db *gorm.DB
}

func (v *VoteRepository) AddVotes(votes *entities.Vote) error {
	err := v.Db.Create(votes).Error
	if err != nil {
		return err
	}
	return nil
}

func (v *VoteRepository) ShowVotes() ([]*entities.Vote, error) {
	var votes []*entities.Vote
	err := v.Db.Find(&votes).Error
	if err != nil {
		return nil, err
	}
	return votes, nil
}

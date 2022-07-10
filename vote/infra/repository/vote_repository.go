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

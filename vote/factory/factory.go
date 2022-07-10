package factory

import (
	"gomq/vote/infra/repository"
	"gomq/vote/usecase"

	"github.com/jinzhu/gorm"
)

func VoteUsecaseFactory(database *gorm.DB) usecase.VoteUseCase {
	voteRepository := repository.VoteRepository{Db: database}
	voteUseCase := usecase.VoteUseCase{
		Irespositoy: &voteRepository,
	}
	return voteUseCase
}

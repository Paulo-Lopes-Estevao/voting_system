package usecase

import "gomq/user/entities"

func UserUsecase(name string) (*entities.User, error) {
	user, err := entities.NewUser(name)
	if err != nil {
		return nil, err
	}
	return user, nil
}

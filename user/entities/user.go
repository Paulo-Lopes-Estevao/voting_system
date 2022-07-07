package entities

import (
	"errors"

	uuid "github.com/satori/go.uuid"
)

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func NewUser(name string) (*User, error) {
	uuid := uuid.NewV4().String()
	user := &User{
		ID:   uuid,
		Name: name,
	}

	err := user.verifyName()

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *User) verifyName() error {
	if u.Name == "" {
		return errors.New("name required")
	}
	return nil
}

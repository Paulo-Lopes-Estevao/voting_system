package entities

import (
	"errors"
	"time"

	uuid "github.com/satori/go.uuid"
)

type User struct {
	ID        string    `json:"id" gorm:"type:uuid;primary_key" valid:"uuid"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at" valid:"-"`
	UpdatedAt time.Time `json:"updated_at" valid:"-"`
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

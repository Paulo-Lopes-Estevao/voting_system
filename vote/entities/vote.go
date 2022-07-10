package entities

import "time"

type Vote struct {
	ID        string    `json:"id" gorm:"type:uuid;primary_key" valid:"uuid"`
	Name      string    `json:"name"`
	Vote      string    `json:"vote"`
	CreatedAt time.Time `json:"created_at" valid:"-"`
	UpdatedAt time.Time `json:"updated_at" valid:"-"`
}

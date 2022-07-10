package dto

import (
	"encoding/json"
	"fmt"

	"github.com/go-playground/validator/v10"
)

type Vote struct {
	ID   string `json:"id" validate:"required,uuid4"`
	Name string `json:"name" validate:"required"`
	Vote string `json:"vote" validate:"required"`
}

func (t *Vote) isValid() error {
	v := validator.New()
	err := v.Struct(t)
	if err != nil {
		fmt.Errorf("Error during voting validation: %s", err.Error())
		return err
	}
	return nil
}

func (t *Vote) ParseJson(data []byte) error {
	err := json.Unmarshal(data, t)
	if err != nil {
		return err
	}

	err = t.isValid()
	if err != nil {
		return err
	}

	return nil
}

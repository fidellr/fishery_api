package utils

import (
	"github.com/go-playground/validator/v10"
)

func Validate(data interface{}) error {
	vld := validator.New()
	err := vld.Struct(data)
	return err
}

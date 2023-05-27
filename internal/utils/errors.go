package utils

import (
	"errors"
	"fmt"
)

var (
	ErrContextNil = errors.New("Context is nil")

	ErrNotFound = errors.New("Your requested item does not exists")
)

type ConstraintError string

func (e ConstraintError) Error() string {
	return string(e)
}

func ConstraintErrorf(format string, a ...interface{}) ConstraintError {
	return ConstraintError(fmt.Sprintf(format, a...))
}

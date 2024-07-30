package utils

import (
	"errors"
)

func FunctionReturnFalseAsError(fn func(input string) bool, message string) func(input string) error {
	return func(newInput string) error {
		if fn(newInput) {
			return nil
		}
		return errors.New(message)
	}
}

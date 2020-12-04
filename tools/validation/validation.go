package validation

import (
	"errors"
)

func ValidateInputNotEmpty(data ...interface{}) error {
	for _, value := range data {
		switch value {
		case "":
			return errors.New("make sure input not empty")
		case 0:
			return errors.New("make sure input not a zero")
		case nil:
			return errors.New("make sure input not a nil")
		}
	}
	return nil
}

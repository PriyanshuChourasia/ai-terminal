package exception

import "errors"

func ErrorNotFound(message string) error {
	if message == "" {
		return errors.New("Not found")
	}
	return errors.New(message)
}

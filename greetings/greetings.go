package greetings

import (
	"errors"
	"fmt"
)

// Hello returns a greeting for the named person

func Hello(name string) (string, error) {
	// if no name given return error with message
	if name == "" {
		return "", errors.New("empty name")
	}

	// if name, return a greeting that embeds name into message
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}

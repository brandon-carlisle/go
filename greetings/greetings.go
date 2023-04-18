package greetings

import "fmt"

// Hello returns a greeting for the named person

func Hello(name string) string {
	// Return a greeting that embeds name into message

	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}

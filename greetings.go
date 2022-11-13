package gogreet

import "fmt"
import "errors"

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	// Return a greeting that embeds the name in a message.
	if name == "" {
		return "", errors.New("empty name")

	}
	message := fmt.Sprintf("Hi, %v. Welcome, LOL!", name)
	return message, nil
}

package gogreet

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	// Return a greeting that embeds the name in a message.
	if name == "" {
		return "", errors.New("empty name")

	}
	message := fmt.Sprintf(randomMessage(), name)
	return message, nil
}

// init sets initial values for variables used in the function.
func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomMessage() string {
	slice := []string{
		"Hi %v",
		"Hello there, %v",
		"What is going on, %v?!",
	}
	return slice[rand.Intn(len(slice))]
}

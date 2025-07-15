package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// Hello return a greeting for the named person

func Hello(name string) string {
	//return a greeting that embeds the name in the message

	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}

func HelloNew(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}

	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func randomFormat() string {
	//A slice of message formats

	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v",
		"Hail, %v! Well met!",
	}

	// returning a randomly selected message format by specifying a random index for the slice of formats.

	return formats[rand.Intn(len(formats))]
}

func Hellos(names []string) (map[string]string, error) {
	// a map to associate names with messages
	messages := make(map[string]string)

	for _, name := range names {
		message, err := HelloNew(name)

		if err != nil {
			return nil, err
		}
		messages[name] = message
	}

	return messages, nil
}

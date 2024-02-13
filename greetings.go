package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("String vacio")
	}
	message := fmt.Sprintf(RandomGreetings(), name)
	return message, nil
}

func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, errors.New("String vacio")
		}
		messages[name] = message
	}
	return messages, nil
}

func RandomGreetings() string {
	formats := []string{
		"Hola %v",
		"Hello %v",
		"Hijo de puta %v",
	}
	return formats[rand.Intn(len(formats))]
}

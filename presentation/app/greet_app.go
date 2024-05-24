package app

import (
	"errors"
	"fmt"
	"strings"
)

type GreetApp struct{}

func NewGreetApp() *GreetApp { return &GreetApp{} }

// Greet returns a greeting for the given name
func (a *GreetApp) Greet(name string) (string, error) {
	name = strings.TrimSpace(name)
	if len(name) == 0 {
		return *new(string), errors.New("name is empty")
	}
	return fmt.Sprintf("Hello %s, It's show time!", name), nil
}

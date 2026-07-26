package repl

import (
	"fmt"
)

func commandHelp(c *config) error {
	commands := getCommands()

	fmt.Printf("Welcome to the Pokedex!\nUsage:\n")
	for _, c := range commands {
		fmt.Printf("%s: %s\n", c.name, c.description)
	}

	return nil
}

package repl

import "fmt"

func commandPokedex(c *config, _ string) error {
	pokedex := c.pokedex.GetAll()

	if len(pokedex) == 0 {
		fmt.Println("You have not caught any pokemon yet, go explore trainer")
	}

	fmt.Println("Your Pokedex:")
	for _, p := range pokedex {
		fmt.Printf(" - %v\n", p.Name)
	}

	return nil
}

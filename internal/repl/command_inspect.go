package repl

import "fmt"

func commandInspect(c *config, name string) error {
	pokemon, isCatched := c.pokedex.Get(name)

	if !isCatched {
		fmt.Println("you have not caught that pokemon")
		return nil
	}

	fmt.Printf("Name: %v\n", pokemon.Name)
	return nil
}

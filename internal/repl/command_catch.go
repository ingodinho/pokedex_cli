package repl

import (
	"fmt"
	"math/rand/v2"
)

func commandCatch(c *config, name string) error {
	if name == "" {
		fmt.Println("please provide a pokemon name")
	}

	fmt.Printf("Throwing a Pokeball at %v...\n", name)

	pokemon, err := c.client.FetchPokemonDetails(name)
	if err != nil {
		return err
	}

	chance := rand.Float32() * 150
	if chance > float32(pokemon.BaseExperience) {
		fmt.Printf("%v was caught!\n", pokemon.Name)
		c.pokedex.Add(pokemon)
	} else {
		fmt.Printf("%v escaped!\n", pokemon.Name)
	}

	return nil
}

package repl

import (
	"fmt"
	"github.com/ingodinho/pokedex_cli/internal/pokeapi"
)

func commandMap(c *config) error {
	client := pokeapi.NewClient()
	locationAreaRes, err := client.FetchLocationAreas(c.Next)
	if err != nil {
		return err
	}

	c.Next = locationAreaRes.Next
	c.Previous = locationAreaRes.Previous

	displayLocationAreas(locationAreaRes.Results)

	return nil
}

func commandMapB(c *config) error {
	client := pokeapi.NewClient()
	if c.Previous == nil {
		fmt.Println("you're on the first page")
		return nil
	}

	locationAreaRes, err := client.FetchLocationAreas(c.Previous)
	if err != nil {
		return err
	}

	c.Next = locationAreaRes.Next
	c.Previous = locationAreaRes.Previous

	displayLocationAreas(locationAreaRes.Results)

	return nil
}

func displayLocationAreas(areas []pokeapi.LocationArea) {
	for _, area := range areas {
		fmt.Println(area.Name)
	}
}

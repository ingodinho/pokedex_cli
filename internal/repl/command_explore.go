package repl

import "fmt"

func commandExplore(c *config, location string) error {
	if location == "" {
		println("please provide a location name")
		return nil
	}

	res, err := c.client.FetchLocationAreaDetails(location)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", location)
	fmt.Println("Found Pokemon:")

	for _, enc := range res.PokemonEncounters {
		fmt.Println(" - ", enc.Pokemon.Name)
	}

	return nil
}

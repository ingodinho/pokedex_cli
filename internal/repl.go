package internal

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	Next string
	Previous string
}

func StartRepl() {
	commands := getCommands()
	c := config{}
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		cleanedInput := cleanInput(input)

		if len(cleanedInput) == 0 {
			continue
		}

		command := cleanedInput[0]

		foundCommand, ok := commands[command]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}

		err := foundCommand.callback(&c)
		if err != nil {
			fmt.Printf("Error happened: %v", err)
			continue
		}

		if scanner.Err() != nil {
			os.Exit(1)
		}
	}
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name: "map",
			description: "Shows the next page of Areas",
			callback: commandMap,
		},
		"mapb": {
			name: "mapb",
			description: "Shows the previous page of Poke Areas",
			callback: commandMapB,
		},
	}
}

func cleanInput(text string) []string {
	lowered := strings.ToLower(text)
	result := strings.Fields(lowered)

	return result
}

func commandExit(c *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(c *config) error {
	commands := getCommands()

	fmt.Printf("Welcome to the Pokedex!\nUsage:\n")
	for _, c := range commands {
		fmt.Printf("%s: %s\n", c.name, c.description)
	}

	return nil
}

func commandMap(c *config) error {
	locationAreaRes, err := fetchLocationAreas(c.Next)
	if err != nil {
		return err
	}

	c.Next = locationAreaRes.Next
	c.Previous = locationAreaRes.Previous

	displayLocationAreas(locationAreaRes.Results)

	return nil
}

func commandMapB(c *config) error {
	if c.Previous == "" {
		fmt.Println("you're on the first page")
		return nil
	}

	locationAreaRes, err := fetchLocationAreas(c.Previous)
	if err != nil {
		return err
	}

	c.Next = locationAreaRes.Next
	c.Previous = locationAreaRes.Previous

	displayLocationAreas(locationAreaRes.Results)

	return nil
}

func displayLocationAreas(areas []LocationArea) {
	for _, area := range areas {
		fmt.Println(area.Name)
	}
}

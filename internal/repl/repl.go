package repl

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/ingodinho/pokedex_cli/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, string) error
}

type config struct {
	Next     *string
	Previous *string
	client   *pokeapi.Client
	pokedex  *pokeapi.Pokedex
}

func StartRepl() {
	pokeClient := pokeapi.NewClient(time.Second * 10)
	pokedex := pokeapi.NewPokedex()
	c := config{
		client:  pokeClient,
		pokedex: pokedex,
	}

	commands := getCommands()
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

		arg := ""
		if len(cleanedInput) > 1 {
			arg = cleanedInput[1]
		}

		err := foundCommand.callback(&c, arg)
		if err != nil {
			fmt.Printf("unhandled error in callback: %v\n", err)
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
			name:        "map",
			description: "Shows the next page of Areas",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Shows the previous page of Poke Areas",
			callback:    commandMapB,
		},
		"explore": {
			name:        "explore",
			description: "Shows the list of possible pokemon encounters for a given Area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Catch a pokemon with the given name",
			callback:    commandCatch,
		},
	}
}

func cleanInput(text string) []string {
	lowered := strings.ToLower(text)
	result := strings.Fields(lowered)

	return result
}

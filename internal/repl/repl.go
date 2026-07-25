package repl

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
	Next *string
	Previous *string
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

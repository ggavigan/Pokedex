package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Welcome to Pokedex! For a list of commands, enter 'help'")
		fmt.Println("Pokedex >")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandWord := words[0]

		command, exists := getCommands()[commandWord]
		if exists {
			err := command.callback()
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("command does not exists")
			continue
		}
	}
}

func cleanInput(text string) []string {
	lower := strings.ToLower(text)
	words := strings.Fields(lower)

	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

type Config struct {
	previousURL string
	nextURL     string
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"mapf": {
			name:        "mapf",
			description: "Next location area page",
			callback:    mapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Previous location area page",
			callback:    mapb,
		},
	}
}

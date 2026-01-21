package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	//"errors"
)

type config struct {
	nextLocationsURL *string
	prevLocationsURL *string
	AreaName         *string
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"explore": {
			name:        "explore",
			description: "Find pokemon in a specific location",
			callback:    commandExplore,
		},
		"mapb": {
			name:        "mapb",
			description: "displays the names of 20 location areas",
			callback:    commandMapb,
		},
		"map": {
			name:        "map",
			description: "displays the names of 20 location areas",
			callback:    commandMap,
		},
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
	}
}

func startfunc(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		my_first_line := cleanInput(scanner.Text())
		if len(my_first_line) == 0 {
			continue
		}
		my_first_word := my_first_line[0]
		if my_first_word == "explore"{ cfg.AreaName = &my_first_line[1]}
		cmd, ok := getCommands()[my_first_word]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}

		err := cmd.callback(cfg)
		if err != nil {
			fmt.Println("Error:", err)
		}
	}
}

func cleanInput(text string) []string {
	var my_string []string

	my_text := strings.ToLower(text)

	my_text = strings.TrimSpace(my_text)

	my_string = strings.Fields(my_text)

	return my_string
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*Config, []string) error
}

var command map[string]cliCommand // declare first

func commandExit(cfg *Config, parsedText []string) error {
	fmt.Println("Closing your pokedex... Goodbye")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *Config, parsedText []string) error {
	for name, cmd := range command {
		fmt.Printf("%s: %s\n", name, cmd.description)
	}
	return nil
}


func init() {
	command = map[string]cliCommand{
		"exit": {
			name:        "Exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "Help",
			description: "Shows commands and descriptions",
			callback:    commandHelp,
		},
		"map": {
			name:        "Map",
			description: "Displays next 20 locations",
			callback:    commandMap,
		},
		"mapb": {
            name:        "Mapb",
            description: "Displays previous 20 locations",
            callback:    commandMapBack, // implement similar to commandMap
    },
		"explore": {
            name:        "Explore",
            description: "Explore locations",
            callback:    commandExplore, 
    },
		"catch": {
            name:        "Catch",
            description: "Try to catch a pokemon",
            callback:    commandCatch, 
    },
		"pokedex": {
            name:        "Pokedex",
            description: "Inspect your pokedex",
            callback:    commandPokedex, 
    },
		"inspect": {
            name:        "Inspect",
            description: "Inspect your pokemon",
            callback:    commandInspect, 
    },
	}
}

func cleanInput(text string) []string {
	text = strings.TrimSpace(text)
	text = strings.ToLower(text) 
	if text == "" {
		return []string{}
	} 
	return strings.Fields(text) 
}

func main() {
	cfg := &Config{}
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("Pokedex >  \n")	
		line, _ := reader.ReadString('\n')
		parsedText := cleanInput(line)	
		inputCmd := parsedText[0]
		cmd, exists := command[inputCmd] 
		if !exists {
			fmt.Println("unknown command: ", inputCmd)
			continue
		}
		err := cmd.callback(cfg, parsedText) 
		if err != nil {
			fmt.Println("Error:", err)
		}
	}
}

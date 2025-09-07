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
	callback    func(*Config, []string, string) error
}

var command map[string]cliCommand // declare first

func commandExit(cfg *Config, parsedText []string, s string) error {
	fmt.Println("Closing your pokedex... Goodbye")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *Config, parsedText []string, s string) error {
	for name, cmd := range command {
		fmt.Printf("%s: %s\n", name, cmd.description)
	}
	return nil
}


func init() {
	command = map[string]cliCommand{
		"exit": {
			name:        "exit",
		description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Shows commands and descriptions",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays next 20 locations",
			callback:    commandMap,
		},
		"mapb": {
            name:        "mapb",
            description: "Displays previous 20 locations",
            callback:    commandMapBack, // implement similar to commandMap
    },
		"explore": {
            name:        "explore",
            description: "Explore locations",
            callback:    commandExplore, 
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
		fmt.Printf("Pokedex >")	
		line, _ := reader.ReadString('\n')
		parsedText := cleanInput(line)	
		inputCmd := parsedText[0]
		cmd, exists := command[inputCmd] 
		if !exists {
			fmt.Println("unknown command: ", inputCmd)
			continue
		}
		args := parsedText[1]
		err := cmd.callback(cfg, parsedText, args) 
		if err != nil {
			fmt.Println("Error:", err)
		}
	}
}

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
	callback    func(*Config) error
}

var command map[string]cliCommand // declare first

func commandExit(cfg *Config) error {
	fmt.Println("Closing your pokedex... Goodbye")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *Config) error {
	for name, cmd := range command {
		fmt.Printf("%s: %s\n", name, cmd.description)
	}
	return nil
}

func commandMap(cfg *Config) error {
	data, err := fetchLocation(cfg.Next)
	if err != nil {
		return fmt.Errorf("error fetchLocation")
	}
	for _, loc := range data.Results {
		fmt.Println(loc.Name)
	}
	cfg.Next = data.Next
	cfg.Previous = data.Previous
	return nil
}
func commandMapBack(cfg *Config) error {
	if cfg.Previous == "" {
		fmt.Println("you're on the first page")
		return nil 
	}
	data, err := fetchLocation(cfg.Previous)
	if err != nil {
		fmt.Println("error fetchLocation")
		return nil
	}
	for _, loc := range data.Results {
		fmt.Println(loc.Name)
	}
	cfg.Next = data.Next
	cfg.Previous = data.Previous
	return  nil
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
		err := cmd.callback(cfg) 
		if err != nil {
			fmt.Println("Error:", err)
		}
	}
}

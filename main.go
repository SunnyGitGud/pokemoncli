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

func commandMap(cfg *Config, parsedText []string) error {
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
func commandMapBack(cfg *Config, parsedText []string) error {
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

func commandExplore(cfg *Config, parsedText []string) error {
	if len(parsedText) < 2 {
		return fmt.Errorf("use Explore + location name")
	}

	locationName := parsedText[1]
	url := "https://pokeapi.co/api/v2/location-area/" + locationName + "/"

	data, err := fetchLocation(url)
	if err != nil {
		return fmt.Errorf("explore fetch error: %v", err)
	}

	for _, loc := range data.Results {
	fmt.Printf("Exploring %s", locationName)
	fmt.Println("-", loc.Name)
	}

	cfg.Next = data.Next
	cfg.Previous = data.Previous

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
		err := cmd.callback(cfg, parsedText) 
		if err != nil {
			fmt.Println("Error:", err)
		}
	}
}

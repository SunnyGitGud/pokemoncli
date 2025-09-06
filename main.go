package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

var command map[string]cliCommand // declare first

func commandExit() error {
	fmt.Println("Closing your pokedex... Goodbye")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	for name, cmd := range command {
		fmt.Printf("%s: %s\n", name, cmd.description)
	}
	return nil
}

func commandMap() error {
	resp, err := http.Get("https://pokeapi.co/api/v2/location-area/")
	if err != nil {
			return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
			return err
	}

	if resp.StatusCode > 299 {
			return fmt.Errorf("Response failed with status %d: %s", resp.StatusCode, string(body))
	}

	fmt.Printf("%s\n", string(body)) 
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
		err := cmd.callback() 
		if err != nil {
			fmt.Println("Error:", err)
		}
	}
}

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)


type URLName struct {
    Name string `json:"name"`
    URL  string `json:"url"`
}

type PokemonEncounter struct {
    Pokemon URLName `json:"pokemon"`
}

type LocationAreaDetail struct {
    Name          string             `json:"name"`
    PokemonList   []PokemonEncounter `json:"pokemon_encounters"`
} 

func FetchLocationDetail(url string)  (*LocationAreaDetail, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error getting location details")
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil ,fmt.Errorf("err reading")
	}

	var data LocationAreaDetail 
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, err
	}

	return &data, nil
}



func commandExplore(cfg *Config, parsedText []string) error {
	if len(parsedText) < 2 {
			return fmt.Errorf("Use Explore <location-name-or-id>")
	}

	// take the *second string* after "explore"
	location := parsedText[1]

	url := "https://pokeapi.co/api/v2/location-area/" + location + "/"
	data, err := FetchLocationDetail(url)
	if err != nil {
			return fmt.Errorf("can't fetch location: %v", err)
	}

	fmt.Printf("Exploring %s:\n", location)
	for _, encounter := range data.PokemonList {
			fmt.Println("-", encounter.Pokemon.Name)
	}

	return nil
}

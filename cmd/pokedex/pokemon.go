package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)


type URLName struct {
	Name string
	URL string
}

type PokemonEncounter struct{
	Pokemon URLName
}

type LocationAreaDetail struct{
	Name string
	PokemonList []PokemonEncounter
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



func commandExplore(cfg *Config, locationName []string) error {
	if len(locationName) < 2 {
		return fmt.Errorf("Use Explore <location-name-or-id>")
	} 
	location := locationName[1]

	url := "https://pokeapi.co/api/v2/location-area/" + location + "/" 
	data, err := FetchLocationDetail(url)
	if err != nil {
		return fmt.Errorf("cant fetch locations") 
	}

	fmt.Printf("Exploring %s: \n", locationName)
	for _, encounter := range data.PokemonList {
		fmt.Println("-", encounter.Pokemon.Name)
	}

	return nil
}

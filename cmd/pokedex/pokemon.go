package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
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

type Pokemon struct {
	ID    int    `json:"id"`
	Name  string `json:"name"` 
	BaseExperience int    `json:"base_experience"`
	Height int   `json:"height"`
	Weight int   `json:"weight"`

	Types []struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"type"`
	} `json:"types"`

	Stats []struct {
		BaseStat int `json:"base_stat"`
		Stat     struct {
			Name string `json:"name"`
		} `json:"stat"`
	} `json:"stats"`

	Sprites struct {
		FrontDefault string `json:"front_default"`
	} `json:"sprites"`
}

var Mypokedex = make(map[string]*Pokemon)	

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

func FetchPokemonDetail(url string) (*Pokemon, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("is not a valid pokemon")
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("err reading")
	}

	var data Pokemon
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, err
	}
	return & data, nil
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


func commandCatch(cfg *Config, parsedText []string) error {
	if len(parsedText) < 2 {
		return fmt.Errorf("Use Catch <pokemon name of this region>")
	}

	pokemonI := parsedText[1]
	url := "https://pokeapi.co/api/v2/pokemon/" + pokemonI + "/" 

	data, err := FetchPokemonDetail(url)
	if err !=nil {
		return fmt.Errorf("%s is not a valid pokemon", pokemonI)
	}	

	fmt.Printf("Throwing a Pokeball at %s...\n", data.Name)
	

	catchChance := 200.0 /float64(data.BaseExperience + 50)
	if catchChance > 0.9 {
		catchChance = 0.9
	}
	if catchChance < 0.1 {
		catchChance= 0.1
	}

	if rand.Float64() < catchChance {
		if _, exists := Mypokedex[data.Name]; exists {
			fmt.Printf("%s is already in your pokedex\n", data.Name)
		} else {
			Mypokedex[data.Name] = data
			fmt.Printf("%s was caught! and added to your pokedex\n", data.Name)	
		}
	} else {
		fmt.Printf("%s escaped\n", data.Name)
	}

	return nil
}	


func commandIpokedex (cfg *Config, parsedText []string) error {
	if len(Mypokedex) == 0 {
		fmt.Println("Your pokedex is Empty go catch some pokemon j")
		return nil
	}

	fmt.Println("your Pokedex: ")
	for name := range Mypokedex{
		fmt.Print("- %s \n", name)
	}
	return nil
}

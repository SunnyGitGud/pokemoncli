package main

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/sunnygitgud/pokemoncli/internal/pokecache"
)

type Config struct {
	Next string
	Previous string
}


type LocationArea struct {
	Name string
	URL string
}

type LocationAreaRespose struct {
	Count int 
	Next string
	Previous string
	Results []LocationArea
} 

var caches =  pokecache.NewCache(5 * time.Second)

func fetchLocation(url string) (*LocationAreaRespose, error){
	if url == "" {
	url = "https://pokeapi.co/api/v2/location-area/"
	}

	if data, ok := caches.Get(url); ok {
		var res LocationAreaRespose
		if err := json.Unmarshal(data, &res); err != nil {
			return nil, err
		}
		return &res, nil
	}

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close() 

	body,err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	caches.Add(url, body)
	var data LocationAreaRespose
	if err := json.Unmarshal(body, &data);err != nil {
		return nil, err
	}
	return &data, nil
}

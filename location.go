package main

import (
	"encoding/json"
	"io"
	"net/http"
)

type Config struct {
	Next string
	Previous string
}

type LocationAreaRespose struct {
	Count int 
	Next string
	Previous string
	Results []LocationArea
} 

type LocationArea struct {
	Name string
	URL string
}



func fetchLocation(url string) (*LocationAreaRespose, error){
	if url == "" {
	url = "https://pokeapi.co/api/v2/location-area/"
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

	var data LocationAreaRespose
	if err := json.Unmarshal(body, &data);err != nil {
		return nil, err
	}

	return &data, nil
}

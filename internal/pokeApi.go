package internal

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type PokemonLocation struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type LocationAreaResponse struct {
	Count    int               `json:"count"`
	Next     string            `json:"next"`
	Previous interface{}       `json:"previous"`
	Results  []PokemonLocation `json:"results"`
}

func GetPokemonMap() *LocationAreaResponse {

	res, err := http.Get("https://pokeapi.co/api/v2/location-area/")
	if err != nil {
		fmt.Printf("Error fetching Pokemon map: %v\n", err)
		return nil
	}
	defer res.Body.Close()

	var locationAreaResponse LocationAreaResponse
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&locationAreaResponse)
	if err != nil {
		fmt.Printf("Error decoding Pokemon map response: %v\n", err)
		return nil
	}

	return &locationAreaResponse
}

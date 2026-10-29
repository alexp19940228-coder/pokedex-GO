package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type pokeapi struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func commandMap(cfg *config) error {
	url := "https://pokeapi.co/api/v2/location-area/"
	if cfg.Next != "" {
		url = cfg.Next
	}
	client := http.Client{}

	res, err := client.Get(url)
	if err != nil {
		return fmt.Errorf("Error getting a HTTP response: %v", err)
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if res.StatusCode > 299 {
		return fmt.Errorf("Response failed with status: %v", res.Status)
	}
	var poke pokeapi

	err = json.Unmarshal(body, &poke)
	if err != nil {
		return fmt.Errorf("Error during Unmarshal: %v", err)
	}
	for _, p := range poke.Results {
		fmt.Println(p.Name)
	}
	cfg.Next = poke.Next
	cfg.Previous = poke.Previous

	return nil
}

func commandMapb(cfg *config) error {
	if cfg.Previous == "" {
		fmt.Println("you're on the first page")
		return nil
	}
	url := cfg.Previous
	client := http.Client{}

	res, err := client.Get(url)
	if err != nil {
		return fmt.Errorf("Error getting a HTTP response: %v", err)
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if res.StatusCode > 299 {
		return fmt.Errorf("Response failed with status: %v", res.Status)
	}
	var poke pokeapi

	err = json.Unmarshal(body, &poke)
	if err != nil {
		return fmt.Errorf("Error during Unmarshal: %v", err)
	}
	for _, p := range poke.Results {
		fmt.Println(p.Name)
	}
	cfg.Next = poke.Next
	cfg.Previous = poke.Previous

	return nil
}

package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/Averagejoestudent/pokedexcli/internal/pokeapi"
)

func commandExit(cfg *config) error { // the exit command
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config) error { // the help command
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}

func commandMap(cfg *config) error { // the Map command
	var url string
	var resp pokeapi.LocationAreaListResponse
	if cfg.nextLocationsURL == nil {
		url = "https://pokeapi.co/api/v2/location-area"
	} else {
		url = *cfg.nextLocationsURL
	}
	body, err := pokeapi.Location(url)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return err
	}
	for i := 0; i < len(resp.Results); i++ {
		fmt.Println(resp.Results[i].Name)
	}
	cfg.nextLocationsURL = resp.Next
	cfg.prevLocationsURL = resp.Previous

	return nil
}

func commandMapb(cfg *config) error { // the Mapb command
	url := *cfg.prevLocationsURL
	var resp pokeapi.LocationAreaListResponse
	if cfg.prevLocationsURL == nil {
		fmt.Println("you're on the first page")
		return nil
	}
	body, err := pokeapi.Location(url)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return err
	}
	for i := 0; i < len(resp.Results); i++ {
		fmt.Println(resp.Results[i].Name)
	}
	cfg.nextLocationsURL = resp.Next
	cfg.prevLocationsURL = resp.Previous

	return nil
}

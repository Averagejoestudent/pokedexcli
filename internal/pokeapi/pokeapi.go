package pokeapi

import (
	"fmt"
	"io"
	"net/http"
)
type BaseExperience struct{
	BaseExp int `json:"base_experience"`
}



type PokemonResponse struct {
	PokemonEncounters []PokemonEncounters `json:"pokemon_encounters"`
}
type PokemonEncounters struct {
	Pokemon Pokemon	`json:"pokemon"`
	} 

type Pokemon struct {
			Name string `json:"name"`
		} 


type NamedAPIResource struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type LocationAreaListResponse struct {
	Count    int                `json:"count"`
	Next     *string            `json:"next"`
	Previous *string            `json:"previous"`
	Results  []NamedAPIResource `json:"results"`
}

func JsonFrmGet(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		return nil, fmt.Errorf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		return nil, err
	}
	return body, nil
}

package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type LocationArea struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous any    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type Config struct {
	Next     string
	Previous string
}

func GetLocationAreas(config *Config) (LocationArea, error) {
	defaultUrl := "https://pokeapi.co/api/v2/location-area"
	if config.Next == "" {
		config.Next = defaultUrl
	}
	url := config.Next
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating a request", err)
		return LocationArea{}, err
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making a request", err)
		return LocationArea{}, err
	}
	defer resp.Body.Close()

	var la LocationArea
	err = json.NewDecoder(resp.Body).Decode(&la)
	if err != nil {
		fmt.Println("Error decoding response body", err)
		return LocationArea{}, err
	}
	config.Next = la.Next
	if la.Previous != nil {
		config.Previous = la.Previous.(string)
	}
	return la, nil
}

func GetLocationAreasBackward(config *Config) (LocationArea, error) {
	defaultUrl := "https://pokeapi.co/api/v2/location-area"
	if config.Previous == "" {
		config.Previous = defaultUrl
	}
	url := config.Previous
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating a request", err)
		return LocationArea{}, err
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making a request", err)
		return LocationArea{}, err
	}
	defer resp.Body.Close()

	var la LocationArea
	err = json.NewDecoder(resp.Body).Decode(&la)
	if err != nil {
		fmt.Println("Error decoding response body", err)
		return LocationArea{}, err
	}
	config.Next = la.Next
	if la.Previous != nil {
		config.Previous = la.Previous.(string)
	}
	return la, nil
}

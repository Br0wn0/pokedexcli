package pokeapi

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func RetrieveData(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln("error", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func ProcessData(url string) (LocationsArea, error) {
	data, err := RetrieveData(url)
	if err != nil {
		return LocationsArea{}, err
	}
	var locations LocationsArea
	err = json.Unmarshal(data, &locations)
	if err != nil {
		return LocationsArea{}, err
	}

	return locations, nil
}

type LocationsArea struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

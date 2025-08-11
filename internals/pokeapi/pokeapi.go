package pokeapi

import (
	"encoding/json"
	"net/http"
)

type areaMap struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Details  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func GetLocations(url string) (locations []string, next, previous string, err error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, "", "", err
	}
	defer response.Body.Close()

	decoder := json.NewDecoder(response.Body)

	area := areaMap{}
	err = decoder.Decode(&area)
	if err != nil {
		return nil, "", "", err
	}

	for _, location := range area.Details {
		locations = append(locations, location.Name)
	}
	next = area.Next
	previous = area.Previous

	return locations, next, previous, nil
}

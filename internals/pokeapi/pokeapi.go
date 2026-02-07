package pokeapi

import (
	"encoding/json"
	"net/http"
)

// location: mt coronet
// area: mt coronet 1f route 207

// locationsMap = general list of areas
// count = how many there are in total
// next = url to the next page of areas
// previous = url to the previous page of areas
// locations = areas
// name = area's name
// url = area's url
type locationsMap struct {
	Count     int    `json:"count"`
	Next      string `json:"next"`
	Previous  string `json:"previous"`
	Locations []struct {
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

	area := locationsMap{}
	err = decoder.Decode(&area)
	if err != nil {
		return nil, "", "", err
	}

	for _, location := range area.Locations {
		locations = append(locations, location.Name)
	}
	next = area.Next
	previous = area.Previous

	return locations, next, previous, nil
}

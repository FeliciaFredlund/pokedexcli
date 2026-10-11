package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/FeliciaFredlund/pokedexcli/internals/pokecache"
)

type Client struct {
	cache      pokecache.Cache
	httpClient http.Client
}

func NewClient(timeout, cacheInterval time.Duration) Client {
	return Client{
		cache: pokecache.NewCache(cacheInterval),
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}

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

func (c *Client) GetLocations(url string) (locations []string, next, previous string, err error) {
	area := locationsMap{}

	if data, exists := c.cache.Get(url); exists {
		err := json.Unmarshal(data, &area)
		if err != nil {
			return nil, "", "", err
		}
	} else {
		response, err := http.Get(url)
		if err != nil {
			return nil, "", "", err
		}
		defer response.Body.Close()

		data, err := io.ReadAll(response.Body)
		if err != nil {
			return nil, "", "", err
		}

		err = json.Unmarshal(data, &area)
		if err != nil {
			return nil, "", "", err
		}

		c.cache.Add(url, data)
	}

	for _, location := range area.Locations {
		locations = append(locations, location.Name)
	}
	next = area.Next
	previous = area.Previous

	return locations, next, previous, nil
}

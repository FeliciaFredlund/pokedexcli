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

// location = general list of areas
// count = how many there are in total
// next = url to the next page of areas
// previous = url to the previous page of areas
// areas
// name = area's name
// url = area's url
type location struct {
	//Count     int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Areas    []struct {
		Name string `json:"name"`
		//URL  string `json:"url"`
	} `json:"results"`
}

func (c *Client) GetLocations(url string) (locations []string, next, previous string, err error) {
	location := location{}

	if data, exists := c.cache.Get(url); exists {
		err := json.Unmarshal(data, &location)
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

		err = json.Unmarshal(data, &location)
		if err != nil {
			return nil, "", "", err
		}

		c.cache.Add(url, data)
	}

	for _, location := range location.Areas {
		locations = append(locations, location.Name)
	}
	next = location.Next
	previous = location.Previous

	return locations, next, previous, nil
}

type area struct {
	//ID                int    `json:"id"`
	//Name              string `json:"name"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			//URL  string `json:"url"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

func (c *Client) GetPokemons(url string) (pokemons []string, err error) {
	area := area{}

	if data, exists := c.cache.Get(url); exists {
		err := json.Unmarshal(data, &area)
		if err != nil {
			return nil, err
		}
	} else {
		response, err := http.Get(url)
		if err != nil {
			return nil, err
		}
		defer response.Body.Close()

		data, err := io.ReadAll(response.Body)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(data, &area)
		if err != nil {
			return nil, err
		}

		c.cache.Add(url, data)
	}

	for _, pokemon := range area.PokemonEncounters {
		pokemons = append(pokemons, pokemon.Pokemon.Name)
	}

	return pokemons, nil
}

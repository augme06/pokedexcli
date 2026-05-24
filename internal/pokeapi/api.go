package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/augme06/pokedexcli/internal/pokecache"
)

// Initial request
const initUrl string = "https://pokeapi.co/api/v2"
const cacheInterval time.Duration = 10 * time.Second

var cache *pokecache.Cache

func GetNextMap(config *Config) error {
	var url string

	switch config.Next {
	case nil:
		url = initUrl + "/location-area" + "?offset=0&limit=20"
	default:
		url = *config.Next
	}

	var data []byte
	if d, ok := Cache(url); ok {
		data = d
	} else {
		d, err := doRequest("GET", url)
		if err != nil {
			return err
		}
		data = d
	}

	cache.Add(url, data)

	info := Response{}
	if err := json.Unmarshal(data, &info); err != nil {
		return err
	}

	config.Next = info.Next
	config.Previous = info.Previous

	for _, c := range info.Results {
		fmt.Println(c.Name)
	}

	return nil
}

func GetPreviousMap(config *Config) error {
	var url string

	switch config.Previous {
	case nil:
		fmt.Println("You are on the first page!")
		return nil
	default:
		url = *config.Previous
	}

	var data []byte
	if d, ok := Cache(url); ok {
		data = d
	} else {
		d, err := doRequest("GET", url)
		if err != nil {
			return err
		}
		data = d
	}

	cache.Add(url, data)

	info := Response{}
	if err := json.Unmarshal(data, &info); err != nil {
		return err
	}

	config.Next = info.Next
	config.Previous = info.Previous

	for _, c := range info.Results {
		fmt.Println(c.Name)
	}

	return nil
}

func Explore(locationArea string) error {
	url := initUrl + "/location-area/" + locationArea

	var data []byte
	if d, ok := Cache(url); ok {
		data = d
	} else {
		d, err := doRequest("GET", url)
		if err != nil {
			return err
		}
		data = d
	}

	cache.Add(url, data)

	info := PokemonResponse{}
	if err := json.Unmarshal(data, &info); err != nil {
		return err
	}

	for _, p := range info.Encounters {
		fmt.Println(p.PokemonRef.Name)
	}

	return nil
}

func GetPokemon(pokemon string) (Pokemon, error) {
	url := initUrl + "/pokemon/" + pokemon

	var data []byte
	if d, ok := Cache(url); ok {
		data = d
	} else {
		d, err := doRequest("GET", url)
		if err != nil {
			return Pokemon{}, err
		}
		data = d
	}

	cache.Add(url, data)

	info := Pokemon{}
	if err := json.Unmarshal(data, &info); err != nil {
		return Pokemon{}, err
	}

	return info, nil
}

func doRequest(method, url string) ([]byte, error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func Cache(key string) ([]byte, bool) {
	if cache == nil {
		cache = pokecache.NewCache(cacheInterval)
	}
	return cache.Get(key)
}

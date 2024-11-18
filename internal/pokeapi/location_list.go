package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (LocList, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	// check the cache
	data, ok := c.cache.Get(url)
	if ok {
		//cache hit
		locList := LocList{}
		err := json.Unmarshal(data, &locList)
		if err != nil {
			return LocList{}, err
		}

		return locList, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocList{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocList{}, err
	}
	defer res.Body.Close()

	data, err = io.ReadAll(res.Body)
	if err != nil {
		return LocList{}, nil
	}

	locList := LocList{}
	err = json.Unmarshal(data, &locList)
	if err != nil {
		return LocList{}, err
	}

	c.cache.Add(url, data)

	return locList, nil
}

func (c *Client) ListPokemonInArea(area *string) (LocationAreaDetails, error) {
	if area == nil {
		return LocationAreaDetails{}, fmt.Errorf("No area to explore was given")
	}
	url := baseURL + "/location-area/" + *area

	// check the cache
	data, ok := c.cache.Get(url)
	if ok {
		//cache hit
		locAreaDetails := LocationAreaDetails{}
		err := json.Unmarshal(data, &locAreaDetails)
		if err != nil {
			return LocationAreaDetails{}, err
		}

		return locAreaDetails, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationAreaDetails{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreaDetails{}, err
	}
	defer res.Body.Close()

	data, err = io.ReadAll(res.Body)
	if err != nil {
		return LocationAreaDetails{}, nil
	}

	locAreaDetails := LocationAreaDetails{}
	err = json.Unmarshal(data, &locAreaDetails)
	if err != nil {
		return LocationAreaDetails{}, err
	}

	c.cache.Add(url, data)

	return locAreaDetails, nil
}

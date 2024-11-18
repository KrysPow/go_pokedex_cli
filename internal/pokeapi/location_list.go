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
		fmt.Println("Cache hit!")
		locList := LocList{}
		err := json.Unmarshal(data, &locList)
		if err != nil {
			return LocList{}, err
		}

		return locList, nil
	}
	fmt.Println("Cache miss!")

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

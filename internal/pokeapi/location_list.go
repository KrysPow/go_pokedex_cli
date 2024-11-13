package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (LocList, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
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

	json_data, err := io.ReadAll(res.Body)
	if err != nil {
		return LocList{}, nil
	}

	locList := LocList{}
	err = json.Unmarshal(json_data, &locList)
	if err != nil {
		return LocList{}, err
	}

	return locList, nil
}

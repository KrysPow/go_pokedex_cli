package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type LocList struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func getLocs(locList *LocList) error {
	var url string
	if locList == nil {
		url = "https://pokeapi.co/api/v2/location-area/"
	} else {
		url = locList.Next
	}
	res, err := http.Get(url)
	if err != nil {
		return err
	}

	json_data, err := io.ReadAll(res.Body)
	defer res.Body.Close()

	err = json.Unmarshal(json_data, &locList)
	if err != nil {
		return err
	}

	for _, location := range locList.Results {
		fmt.Println(location.Name)
	}
	return nil
}

func main() {
	var ll *LocList
	getLocs(ll.Next)
}

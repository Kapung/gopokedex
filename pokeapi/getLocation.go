package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (client *Client) PrintLocations(url *string) (Location, error) {
	endpoint := "/location-area"
	getURL := baseURL + endpoint

	if url != nil {
		getURL = *url
	}

	data, ok := client.cache.Get(getURL)

	if ok {
		locations := Location{}
		err := json.Unmarshal(data, &locations)

		if err != nil {
			return Location{}, err
		}

		return locations, nil
	}

	request, err := http.NewRequest("GET", getURL, nil)
	if err != nil {
		return Location{}, err
	}

	response, err := client.httpClient.Do(request)
	if err != nil {
		return Location{}, err
	}

	defer response.Body.Close()

	data, err = io.ReadAll(response.Body)
	if err != nil {
		return Location{}, err
	}

	locations := Location{}
	err = json.Unmarshal(data, &locations)

	if err != nil {
		return Location{}, err
	}

	client.cache.Add(getURL, data)

	return locations, nil
}

func (client *Client) PrintRoute(name string) (Route, error) {
	endpoint := "/location-area/" + name
	getURL := baseURL + endpoint

	data, ok := client.cache.Get(getURL)

	if ok {
		locations := Route{}
		err := json.Unmarshal(data, &locations)

		if err != nil {
			return Route{}, err
		}

		return Route{}, nil
	}

	request, err := http.NewRequest("GET", getURL, nil)
	if err != nil {
		return Route{}, err
	}

	response, err := client.httpClient.Do(request)
	if err != nil {
		return Route{}, err
	}

	defer response.Body.Close()

	data, err = io.ReadAll(response.Body)
	if err != nil {
		return Route{}, err
	}

	route := Route{}
	err = json.Unmarshal(data, &route)

	if err != nil {
		return Route{}, err
	}

	client.cache.Add(getURL, data)

	return route, nil
}

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

package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (client *Client) PrintLocations() (Location, error) {
	endpoint := "/location-area"
	getURL := baseURL + endpoint

	request, err := http.NewRequest("GET", getURL, nil)
	if err != nil {
		return Location{}, err
	}

	response, err := client.httpClient.Do(request)
	if err != nil {
		return Location{}, err
	}

	defer response.Body.Close()

	data, err := io.ReadAll(response.Body)
	if err != nil {
		return Location{}, err
	}

	locations := Location{}
	err = json.Unmarshal(data, &locations)

	if err != nil {
		return Location{}, err
	}

	return locations, nil
}

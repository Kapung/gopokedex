package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (client *Client) GetPokemon(name string) (Pokemon, error) {
	endpoint := "/pokemon/" + name
	getURL := baseURL + endpoint

	data, ok := client.cache.Get(getURL)

	if ok {
		pokemon := Pokemon{}
		err := json.Unmarshal(data, &pokemon)

		if err != nil {
			return Pokemon{}, err
		}

		return pokemon, nil
	}

	request, err := http.NewRequest("GET", getURL, nil)
	if err != nil {
		return Pokemon{}, err
	}

	response, err := client.httpClient.Do(request)
	if err != nil {
		return Pokemon{}, err
	}

	defer response.Body.Close()

	data, err = io.ReadAll(response.Body)
	if err != nil {
		return Pokemon{}, err
	}

	pokemon := Pokemon{}
	err = json.Unmarshal(data, &pokemon)

	if err != nil {
		return Pokemon{}, err
	}

	client.cache.Add(getURL, data)

	return pokemon, nil
}

package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocation(pageUrl *string) (LocationArea, error) {

	url := baseApiUrl + "/location-area"

	if pageUrl != nil {
		url = *pageUrl
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationArea{}, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationArea{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationArea{}, err

	}

	locationResp := LocationArea{}
	err = json.Unmarshal(dat, &locationResp)
	if err != nil {
		return LocationArea{}, err
	}

	return locationResp, nil
}

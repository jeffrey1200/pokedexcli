package pokeapi

// import (
// 	"encoding/json"
// 	"fmt"
// 	"io"
// 	"net/http"
// )

// type LocationArea struct {
// 	Count    int     `json:"count"`
// 	Next     *string `json:"next"`
// 	Previous *string `json:"previous"`
// 	Results  []struct {
// 		Name string `json:"name"`
// 		URL  string `json:"url"`
// 	} `json:"results"`
// }

const (
	baseApiUrl = "https://pokeapi.co/api/v2"
)

// func GetLocationAreas(offset, limit int) (*LocationArea, error) {

// 	res, err := http.Get(fmt.Sprintf("%s?offset=%d&limit=%d", baseApiUrl, offset, limit))
// 	if err != nil {
// 		return nil, err
// 	}

// 	body, err := io.ReadAll(res.Body)
// 	defer res.Body.Close()

// 	if res.StatusCode > 299 {
// 		return nil, fmt.Errorf("response failed with status code: %d and body: %s", res.StatusCode, body)

// 	}
// 	if err != nil {
// 		return nil, err
// 	}
// 	var locAreas LocationArea
// 	err = json.Unmarshal(body, &locAreas)
// 	if err != nil {
// 		return nil, err
// 	}
// 	// t := *pokemonAreaList.Next
// 	// fmt.Print(t)
// 	// for _, areas := range pokemonAreaList.Results {

// 	// 	fmt.Printf("%s\n", areas.Name)
// 	// }
// 	return &locAreas, nil
// 	// fmt.Printf("%s", body)
// }

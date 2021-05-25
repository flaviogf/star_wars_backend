package planets

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
)

type HttpService struct {
}

func (s HttpService) CountMovies(ctx context.Context, planetName string) (int, error) {
	planets, err := getPlanets("https://swapi.dev/api/planets/?page=1", make([]planet, 0))

	if err != nil {
		return 0, err
	}

	for _, it := range planets {
		if it.Name == planetName {
			return len(it.Movies), nil
		}
	}

	return 0, nil
}

type body struct {
	Next    *string  `json:"next"`
	Results []planet `json:"results"`
}

type planet struct {
	Name   string   `json:"name"`
	Movies []string `json:"films"`
}

func getPlanets(url string, planets []planet) ([]planet, error) {
	if len(url) == 0 {
		return planets, nil
	}

	resp, err := http.Get(url)

	if err != nil {
		log.Println(err)

		return []planet{}, err
	}

	var data body

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		log.Println(err)

		return []planet{}, err
	}

	planets = append(planets, data.Results...)

	if data.Next != nil {
		url = *data.Next
	} else {
		url = ""
	}

	return getPlanets(url, planets)
}

package planets

import (
	"context"
	"errors"
)

var (
	ErrPlanetsNotFound = errors.New("planets not found")
)

type ListPlanetHandler struct {
	repository Repository
}

type ListPlanetResponse []PlanetResponse

type PlanetResponse struct {
	ID      interface{} `json:"id"`
	Name    string      `json:"name"`
	Climate string      `json:"climate"`
	Terrain string      `json:"terrain"`
	Movies  int         `json:"movies"`
}

func NewListPlanetHandler(repository Repository) ListPlanetHandler {
	return ListPlanetHandler{
		repository: repository,
	}
}

func (h *ListPlanetHandler) Execute(ctx context.Context) (ListPlanetResponse, error) {
	planets, err := h.repository.GetAll(ctx)

	if err != nil {
		return ListPlanetResponse{}, ErrPlanetsNotFound
	}

	result := ListPlanetResponse{}

	for _, planet := range planets {
		result = append(result, PlanetResponse(planet))
	}

	return result, nil
}

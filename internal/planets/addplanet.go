package planets

import (
	"context"
	"errors"
)

var (
	ErrPlanetNotAdded = errors.New("could not save the planet")
)

type AddPlanetHandler struct {
	repository Repository
}

type AddPlanetRequest struct {
	Name    string `json:"name"`
	Climate string `json:"climate"`
	Terrain string `json:"terrain"`
}

func NewAddPlanetHandler(repository Repository) AddPlanetHandler {
	return AddPlanetHandler{
		repository: repository,
	}
}

func NewAddPlanetRequest(name, climate, terrain string) AddPlanetRequest {
	return AddPlanetRequest{
		Name:    name,
		Climate: climate,
		Terrain: terrain,
	}
}

func (h AddPlanetHandler) Execute(ctx context.Context, request AddPlanetRequest) error {
	planet := NewPlanet(request.Name, request.Climate, request.Terrain)

	return h.repository.Add(ctx, planet)
}

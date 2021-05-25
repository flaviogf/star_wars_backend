package planets

import "context"

type AddPlanetHandler struct {
	repository Repository
	service    Service
}

type AddPlanetRequest struct {
	Name    string `json:"name"`
	Climate string `json:"climate"`
	Terrain string `json:"terrain"`
}

func NewAddPlanetHandler(repository Repository, service Service) AddPlanetHandler {
	return AddPlanetHandler{
		repository: repository,
		service:    service,
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
	if h.repository.Exists(ctx, request.Name) {
		return ErrPlanetAlreadyRegistered
	}

	movies, err := h.service.CountMovies(ctx, request.Name)

	if err != nil {
		return ErrPlanetNotFound
	}

	planet := NewPlanet(request.Name, request.Climate, request.Terrain, movies)

	err = h.repository.Add(ctx, planet)

	if err != nil {
		return ErrPlanetNotAdded
	}

	return nil
}

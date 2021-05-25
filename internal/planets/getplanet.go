package planets

import "context"

type GetPlanetHandler struct {
	repository Repository
}

func NewGetPlanetHandler(repository Repository) GetPlanetHandler {
	return GetPlanetHandler{
		repository: repository,
	}
}

func (h GetPlanetHandler) Execute(ctx context.Context, id interface{}) (GetPlanetResponse, error) {
	planet, err := h.repository.Get(ctx, id)

	if err != nil {
		return GetPlanetResponse{}, ErrPlanetNotFound
	}

	return GetPlanetResponse(planet), nil
}

package planets

import "context"

type GetPlanetByNameHandler struct {
	repositry Repository
}

func NewGetPlanetByNameHandler(repository Repository) GetPlanetByNameHandler {
	return GetPlanetByNameHandler{
		repositry: repository,
	}
}

func (h GetPlanetByNameHandler) Execute(ctx context.Context, name string) (GetPlanetResponse, error) {
	planet, err := h.repositry.GetByName(ctx, name)

	if err != nil {
		return GetPlanetResponse{}, ErrPlanetNotFound
	}

	return GetPlanetResponse(planet), nil
}

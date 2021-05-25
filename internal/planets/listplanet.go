package planets

import "context"

type ListPlanetHandler struct {
	repository Repository
}

type ListPlanetResponse []GetPlanetResponse

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
		result = append(result, GetPlanetResponse(planet))
	}

	return result, nil
}

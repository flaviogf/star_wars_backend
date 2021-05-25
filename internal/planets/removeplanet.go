package planets

import "context"

type RemovePlanetHandler struct {
	repository Repository
}

func NewRemovePlanetHandler(repository Repository) RemovePlanetHandler {
	return RemovePlanetHandler{
		repository: repository,
	}
}

func (h RemovePlanetHandler) Execute(ctx context.Context, id interface{}) error {
	err := h.repository.Remove(ctx, id)

	if err != nil {
		return ErrPlanetNotRemoved
	}

	return nil
}

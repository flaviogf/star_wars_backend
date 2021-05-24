package planets

import (
	"context"
	"testing"
)

func TestNewPlanetShouldReturnNil(t *testing.T) {
	repository := SuccessRepository{}

	sut := NewAddPlanetHandler(repository)

	request := NewAddPlanetRequest("Tatooine", "Hot", "Desert")

	err := sut.Execute(context.Background(), request)

	if err != nil {
		t.Errorf("Expected: %v, Got: %v", nil, err)
	}
}

func TestNewPlanetShouldReturnErrPlanetNotAddedWhenRepositoryFails(t *testing.T) {
	repository := FailureRepository{}

	sut := NewAddPlanetHandler(repository)

	request := NewAddPlanetRequest("Tatooine", "Hot", "Desert")

	err := sut.Execute(context.Background(), request)

	if err != ErrPlanetNotAdded {
		t.Errorf("Expected: %v, Got: %v", ErrPlanetNotAdded, err)
	}
}

type SuccessRepository struct {
}

func (r SuccessRepository) Add(ctx context.Context, planet Planet) error {
	return nil
}

type FailureRepository struct {
}

func (r FailureRepository) Add(ctx context.Context, planet Planet) error {
	return ErrPlanetNotAdded
}

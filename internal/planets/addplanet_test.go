package planets

import (
	"context"
	"errors"
	"testing"
)

func TestAddPlanetShouldReturnNil(t *testing.T) {
	repository := SuccessRepository{}

	service := SuccessService{}

	sut := NewAddPlanetHandler(repository, service)

	request := NewAddPlanetRequest("Tatooine", "Hot", "Desert")

	err := sut.Execute(context.Background(), request)

	if err != nil {
		t.Errorf("Expected: %v, Got: %v", nil, err)
	}
}

func TestAddPlanetShouldReturnErrPlanetNotAddedWhenRepositoryFails(t *testing.T) {
	repository := FailureRepository{}

	service := SuccessService{}

	sut := NewAddPlanetHandler(repository, service)

	request := NewAddPlanetRequest("Tatooine", "Hot", "Desert")

	err := sut.Execute(context.Background(), request)

	if err != ErrPlanetNotAdded {
		t.Errorf("Expected: %v, Got: %v", ErrPlanetNotAdded, err)
	}
}

func TestAddPlanetShouldReturnErrPlanetNotFoundWhenServiceFails(t *testing.T) {
	repository := SuccessRepository{}

	service := FailureService{}

	sut := NewAddPlanetHandler(repository, service)

	request := NewAddPlanetRequest("Wrong", "Hot", "Desert")

	err := sut.Execute(context.Background(), request)

	if err != ErrPlanetNotFound {
		t.Errorf("Expected: %v, Got: %v", ErrPlanetNotFound, err)
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
	return errors.New("something goes wrong")
}

type SuccessService struct {
}

func (s SuccessService) CountMovies(ctx context.Context, planetName string) (int, error) {
	return 0, nil
}

type FailureService struct {
}

func (s FailureService) CountMovies(ctx context.Context, planetName string) (int, error) {
	return 0, errors.New("something goes wrong")
}

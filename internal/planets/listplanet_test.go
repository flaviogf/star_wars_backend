package planets

import (
	"context"
	"testing"
)

func TestListPlanetShouldReturnListPlanetResponseWithOnePlanet(t *testing.T) {
	repository := SuccessRepository{}

	sut := NewListPlanetHandler(repository)

	response, _ := sut.Execute(context.Background())

	if len(response) == 0 {
		t.Errorf("Expected: %v, Got: %v", 1, len(response))
	}
}

func TestListPlanetShouldReturnErrPlanetsNotFoundWhenRepositoryFails(t *testing.T) {
	repository := FailureRepository{}

	sut := NewListPlanetHandler(repository)

	_, err := sut.Execute(context.Background())

	if err != ErrPlanetsNotFound {
		t.Errorf("Expected: %v, Got: %v", err, ErrPlanetsNotFound)
	}
}

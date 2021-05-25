package planets

import (
	"context"
	"testing"
)

func TestGetPlanetShouldReturnPlanet(t *testing.T) {
	repository := SuccessRepository{}

	sut := NewGetPlanetHandler(repository)

	id := "XXXX"

	planet, _ := sut.Execute(context.Background(), id)

	if planet.ID != id {
		t.Errorf("Expected: %v, Got: %v", id, planet.ID)
	}
}

func TestGetPlanetShouldReturnErrPlanetNotFoundWhenRepositoryFails(t *testing.T) {
	repository := FailureRepository{}

	sut := NewGetPlanetHandler(repository)

	id := "XXXX"

	_, err := sut.Execute(context.Background(), id)

	if err != ErrPlanetNotFound {
		t.Errorf("Expected: %v, Got: %v", ErrPlanetNotFound, err)
	}
}

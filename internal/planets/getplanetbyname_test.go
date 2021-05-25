package planets

import (
	"context"
	"testing"
)

func TestGetPlanetByNameShouldReturnPlanet(t *testing.T) {
	repository := SuccessRepository{}

	sut := NewGetPlanetByNameHandler(repository)

	name := "Tatooine"

	planet, _ := sut.Execute(context.Background(), name)

	if planet.Name != name {
		t.Errorf("Expected: %v, Got: %v", name, planet.Name)
	}
}

func TestGetPlanetByNameShouldReturnErrPlanetNotFoundWhenRepositoryFails(t *testing.T) {
	repository := FailureRepository{}

	sut := NewGetPlanetByNameHandler(repository)

	name := "Tatooine"

	_, err := sut.Execute(context.Background(), name)

	if err != ErrPlanetNotFound {
		t.Errorf("Expected: %v, Got: %v", ErrPlanetNotFound, err)
	}
}

package planets

import (
	"context"
	"testing"
)

func TestRemovePlanetShouldReturnNil(t *testing.T) {
	repository := SuccessRepository{}

	sut := NewRemovePlanetHandler(repository)

	id := "XXXX"

	err := sut.Execute(context.Background(), id)

	if err != nil {
		t.Errorf("Expected: %v, Got: %v", nil, err)
	}
}

func TestRemovePlanetShouldReturnErrPlanetNotRemovedWhenRepositoryFails(t *testing.T) {
	repository := FailureRepository{}

	sut := NewRemovePlanetHandler(repository)

	id := "XXXX"

	err := sut.Execute(context.Background(), id)

	if err != ErrPlanetNotRemoved {
		t.Errorf("Expected: %v, Got: %v", ErrPlanetNotRemoved, err)
	}
}

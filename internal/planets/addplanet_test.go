package planets

import (
	"context"
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

func TestAddPlanetShouldReturnErrPlanetAlreadyRegisteredWhenPlanetExists(t *testing.T) {
	repository := PlanetAlreadyRegisteredRepository{}

	service := SuccessService{}

	sut := NewAddPlanetHandler(repository, service)

	request := NewAddPlanetRequest("Wrong", "Hot", "Desert")

	err := sut.Execute(context.Background(), request)

	if err != ErrPlanetAlreadyRegistered {
		t.Errorf("Expected: %v, Got: %v", ErrPlanetAlreadyRegistered, err)
	}
}

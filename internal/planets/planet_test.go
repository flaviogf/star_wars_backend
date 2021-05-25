package planets

import (
	"context"
	"errors"
)

type SuccessRepository struct {
}

func (r SuccessRepository) Add(ctx context.Context, planet Planet) error {
	return nil
}

func (r SuccessRepository) GetAll(ctx context.Context) ([]Planet, error) {
	return []Planet{{"XXXX", "Tatooine", "Hot", "Desert", 5}}, nil
}

func (r SuccessRepository) Get(ctx context.Context, id interface{}) (Planet, error) {
	return Planet{"XXXX", "Tatooine", "Hot", "Desert", 5}, nil
}

func (r SuccessRepository) Remove(ctx context.Context, id interface{}) error {
	return nil
}

type FailureRepository struct {
}

func (r FailureRepository) Add(ctx context.Context, planet Planet) error {
	return errors.New("something goes wrong")
}

func (r FailureRepository) GetAll(ctx context.Context) ([]Planet, error) {
	return []Planet{}, errors.New("something goes wrong")
}

func (r FailureRepository) Get(ctx context.Context, id interface{}) (Planet, error) {
	return Planet{}, errors.New("something goes wrong")
}

func (r FailureRepository) Remove(ctx context.Context, id interface{}) error {
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

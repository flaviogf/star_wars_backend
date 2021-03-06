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

func (r SuccessRepository) GetByName(ctx context.Context, name string) (Planet, error) {
	return Planet{"XXXX", "Tatooine", "Hot", "Desert", 5}, nil
}

func (r SuccessRepository) Remove(ctx context.Context, id interface{}) error {
	return nil
}

func (r SuccessRepository) Exists(ctx context.Context, name string) bool {
	return false
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

func (r FailureRepository) GetByName(ctx context.Context, name string) (Planet, error) {
	return Planet{}, errors.New("something goes wrong")
}

func (r FailureRepository) Remove(ctx context.Context, id interface{}) error {
	return errors.New("something goes wrong")
}

func (r FailureRepository) Exists(ctx context.Context, name string) bool {
	return false
}

type PlanetAlreadyRegisteredRepository struct{}

func (r PlanetAlreadyRegisteredRepository) Add(ctx context.Context, planet Planet) error {
	return errors.New("something goes wrong")
}

func (r PlanetAlreadyRegisteredRepository) GetAll(ctx context.Context) ([]Planet, error) {
	return []Planet{}, errors.New("something goes wrong")
}

func (r PlanetAlreadyRegisteredRepository) Get(ctx context.Context, id interface{}) (Planet, error) {
	return Planet{}, errors.New("something goes wrong")
}

func (r PlanetAlreadyRegisteredRepository) GetByName(ctx context.Context, name string) (Planet, error) {
	return Planet{}, errors.New("something goes wrong")
}

func (r PlanetAlreadyRegisteredRepository) Remove(ctx context.Context, id interface{}) error {
	return errors.New("something goes wrong")
}

func (r PlanetAlreadyRegisteredRepository) Exists(ctx context.Context, name string) bool {
	return true
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

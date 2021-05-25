package planets

import "context"

type Repository interface {
	Add(ctx context.Context, planet Planet) error
	GetAll(ctx context.Context) ([]Planet, error)
	Get(ctx context.Context, id interface{}) (Planet, error)
	GetByName(ctx context.Context, name string) (Planet, error)
	Remove(ctx context.Context, id interface{}) error
}

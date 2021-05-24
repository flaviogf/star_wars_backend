package planets

import "context"

type Repository interface {
	Add(ctx context.Context, planet Planet) error
}

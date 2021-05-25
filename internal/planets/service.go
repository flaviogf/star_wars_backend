package planets

import "context"

type Service interface {
	CountMovies(ctx context.Context, planetName string) (int, error)
}

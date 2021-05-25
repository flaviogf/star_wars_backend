package planets

import (
	"context"

	"github.com/flaviogf/star_wars_backend/cmd/server/database"
	"github.com/flaviogf/star_wars_backend/internal/planets"
	"go.mongodb.org/mongo-driver/bson"
)

type MongoRepository struct{}

func (r MongoRepository) Add(ctx context.Context, planet planets.Planet) error {
	res, err := database.Conn.Collection("planets").InsertOne(ctx, planet)

	if err != nil {
		return err
	}

	planet.ID = res.InsertedID

	return nil
}

func (r MongoRepository) GetAll(ctx context.Context) ([]planets.Planet, error) {
	cur, err := database.Conn.Collection("planets").Find(ctx, bson.D{})

	if err != nil {
		return []planets.Planet{}, err
	}

	defer cur.Close(ctx)

	result := make([]planets.Planet, 0)

	for cur.Next(ctx) {
		var item planets.Planet

		if err = cur.Decode(&item); err != nil {
			return []planets.Planet{}, err
		}

		result = append(result, item)
	}

	return result, nil
}

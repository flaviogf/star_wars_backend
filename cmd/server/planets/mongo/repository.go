package mongo

import (
	"context"
	"log"

	"github.com/flaviogf/star_wars_backend/cmd/server/database"
	"github.com/flaviogf/star_wars_backend/internal/planets"
	"go.mongodb.org/mongo-driver/bson"
)

type MongoRepository struct{}

func (r MongoRepository) Add(ctx context.Context, planet planets.Planet) error {
	res, err := database.Conn.Collection("planets").InsertOne(ctx, bson.D{
		{"name", planet.Name},
		{"climate", planet.Climate},
		{"terrain", planet.Terrain},
	})

	if err != nil {
		log.Println(err)

		return planets.ErrPlanetNotAdded
	}

	planet.ID = res.InsertedID

	return nil
}

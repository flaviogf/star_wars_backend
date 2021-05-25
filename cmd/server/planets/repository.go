package planets

import (
	"context"
	"fmt"

	"github.com/flaviogf/star_wars_backend/cmd/server/database"
	"github.com/flaviogf/star_wars_backend/internal/planets"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (r MongoRepository) Get(ctx context.Context, id interface{}) (planets.Planet, error) {
	objectID, _ := primitive.ObjectIDFromHex((fmt.Sprintf("%v", id)))

	res := database.Conn.Collection("planets").FindOne(ctx, bson.D{{"_id", objectID}})

	if err := res.Err(); err != nil {
		return planets.Planet{}, err
	}

	var planet planets.Planet

	if err := res.Decode(&planet); err != nil {
		return planets.Planet{}, err
	}

	return planet, nil
}

func (r MongoRepository) GetByName(ctx context.Context, name string) (planets.Planet, error) {
	res := database.Conn.Collection("planets").FindOne(ctx, bson.D{{"name", name}})

	if err := res.Err(); err != nil {
		return planets.Planet{}, err
	}

	var planet planets.Planet

	if err := res.Decode(&planet); err != nil {
		return planets.Planet{}, err
	}

	return planet, nil
}

func (r MongoRepository) Remove(ctx context.Context, id interface{}) error {
	objectID, _ := primitive.ObjectIDFromHex((fmt.Sprintf("%v", id)))

	res := database.Conn.Collection("planets").FindOneAndDelete(ctx, bson.D{{"_id", objectID}})

	if err := res.Err(); err != nil {
		return err
	}

	return nil
}

func (r MongoRepository) Exists(ctx context.Context, name string) bool {
	res := database.Conn.Collection("planets").FindOne(ctx, bson.D{{"name", name}})

	if err := res.Err(); err != nil {
		return false
	}

	var planet planets.Planet

	if err := res.Decode(&planet); err != nil {
		return false
	}

	return true
}

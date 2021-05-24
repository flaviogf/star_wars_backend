package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/flaviogf/star_wars_backend/cmd/server/database"
	"github.com/flaviogf/star_wars_backend/cmd/server/routes"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("DATABASE_URL")))

	if err != nil {
		log.Fatal(err)
	}

	database.Conn = client.Database("star_wars")

	http.Handle("/", routes.NewHandler())

	log.Println(http.ListenAndServe(":80", nil))
}

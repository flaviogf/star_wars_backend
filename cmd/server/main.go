package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/flaviogf/star_wars_backend/cmd/server/database"
	"github.com/flaviogf/star_wars_backend/cmd/server/routes"
	"github.com/go-openapi/runtime/middleware"
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

	http.Handle("/docs", middleware.SwaggerUI(middleware.SwaggerUIOpts{SpecURL: "swagger.yaml"}, nil))

	http.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))

	log.Println(http.ListenAndServe(":80", nil))
}

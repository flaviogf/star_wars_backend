package routes

import (
	"net/http"

	"github.com/flaviogf/star_wars_backend/cmd/server/planets"
	"github.com/gorilla/mux"
)

func NewHandler() http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/planets", planets.AddPlanetHandler)

	return r
}

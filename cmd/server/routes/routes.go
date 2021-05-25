package routes

import (
	"net/http"

	"github.com/flaviogf/star_wars_backend/cmd/server/middlewares"
	"github.com/flaviogf/star_wars_backend/cmd/server/planets"
	"github.com/gorilla/mux"
)

func NewHandler() http.Handler {
	r := mux.NewRouter()

	r.Use(middlewares.Json)

	r.HandleFunc("/planets", planets.AddPlanetHandler).Methods(http.MethodPost)

	r.HandleFunc("/planets", planets.ListPlanetHandler).Methods(http.MethodGet)

	r.HandleFunc("/planets/{id}", planets.GetPlanetHandler).Methods(http.MethodGet)

	r.HandleFunc("/planets/{id}", planets.RemovePlanetHandler).Methods(http.MethodDelete)

	return r
}

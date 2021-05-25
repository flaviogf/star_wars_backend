package planets

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/flaviogf/star_wars_backend/internal/planets"
	"github.com/gorilla/mux"
)

func GetPlanetHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	handler := planets.NewGetPlanetHandler(MongoRepository{})

	response, err := handler.Execute(r.Context(), vars["id"])

	if err != nil && (err == planets.ErrPlanetNotFound) {
		log.Println(err)

		w.WriteHeader(http.StatusNotFound)

		return
	}

	if err != nil {
		log.Println(err)

		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	json.NewEncoder(w).Encode(response)
}

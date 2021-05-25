package planets

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/flaviogf/star_wars_backend/internal/planets"
)

func ListPlanetHandler(w http.ResponseWriter, r *http.Request) {
	handler := planets.NewListPlanetHandler(MongoRepository{})

	response, err := handler.Execute(r.Context())

	if err != nil {
		log.Println(err)

		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	json.NewEncoder(w).Encode(response)
}

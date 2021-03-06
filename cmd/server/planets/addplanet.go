package planets

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/flaviogf/star_wars_backend/internal/planets"
)

func AddPlanetHandler(w http.ResponseWriter, r *http.Request) {
	var request planets.AddPlanetRequest

	dec := json.NewDecoder(r.Body)

	if err := dec.Decode(&request); err != nil {
		log.Println(err)

		w.WriteHeader(http.StatusUnprocessableEntity)

		return
	}

	handler := planets.NewAddPlanetHandler(MongoRepository{}, HttpService{})

	err := handler.Execute(r.Context(), request)

	if err != nil && (err == planets.ErrPlanetNotAdded || err == planets.ErrPlanetNotFound || err == planets.ErrPlanetAlreadyRegistered) {
		log.Println(err)

		w.WriteHeader(http.StatusUnprocessableEntity)

		return
	}

	if err != nil {
		log.Println(err)

		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	w.WriteHeader(http.StatusCreated)
}

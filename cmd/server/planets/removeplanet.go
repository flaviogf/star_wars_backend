package planets

import (
	"log"
	"net/http"

	"github.com/flaviogf/star_wars_backend/internal/planets"
	"github.com/gorilla/mux"
)

func RemovePlanetHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	handler := planets.NewRemovePlanetHandler(MongoRepository{})

	err := handler.Execute(r.Context(), vars["id"])

	if err != nil && err == planets.ErrPlanetNotRemoved {
		log.Println(err)

		w.WriteHeader(http.StatusNotFound)

		return
	}

	if err != nil {
		log.Println(err)

		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	w.WriteHeader(http.StatusNoContent)
}

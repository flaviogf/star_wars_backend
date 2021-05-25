package planets

import "errors"

var (
	ErrPlanetNotAdded  = errors.New("could not save the planet")
	ErrPlanetsNotFound = errors.New("planets not found")
	ErrPlanetNotFound  = errors.New("planet not found")
)

type Planet struct {
	ID      interface{} `bson:"_id,omitempty"`
	Name    string      `bson:"name"`
	Climate string      `bson:"climate"`
	Terrain string      `bson:"terrain"`
	Movies  int         `bson:"movies"`
}

func NewPlanet(name, climate, terrain string, movies int) Planet {
	return Planet{
		Name:    name,
		Climate: climate,
		Terrain: terrain,
		Movies:  movies,
	}
}

type GetPlanetResponse struct {
	ID      interface{} `json:"id"`
	Name    string      `json:"name"`
	Climate string      `json:"climate"`
	Terrain string      `json:"terrain"`
	Movies  int         `json:"movies"`
}

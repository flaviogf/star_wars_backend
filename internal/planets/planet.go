package planets

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

package planets

type Planet struct {
	ID      interface{}
	Name    string
	Climate string
	Terrain string
	Movies  int
}

func NewPlanet(name, climate, terrain string, movies int) Planet {
	return Planet{
		Name:    name,
		Climate: climate,
		Terrain: terrain,
		Movies:  movies,
	}
}

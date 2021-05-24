package planets

type Planet struct {
	ID      interface{}
	Name    string
	Climate string
	Terrain string
}

func NewPlanet(name, climate, terrain string) Planet {
	return Planet{
		Name:    name,
		Climate: climate,
		Terrain: terrain,
	}
}

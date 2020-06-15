package planets

// Repository is the repository interface for a planet.
type Repository interface {
	List() ([]Planet, error)
	Exists(name string) (bool, error)
	Insert(planet Planet) error
	FindByName(name string) (Planet, error)
	Delete(name string) error
}

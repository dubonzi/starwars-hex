package planets

// Service is a service interface for planets.
type Service interface {
	List() ([]Planet, error)
	FindByID(id string) (Planet, error)
	FindByName(name string) (Planet, error)
	Add(planet Planet) (Planet, error)
	Delete(id string) error
}

type planetService struct {
	repo Repository
}

// NewService creates a new planets Service.
func NewService(repo Repository) Service {
	return planetService{repo}
}

func (planetService) List() ([]Planet, error) {
	panic("not implemented") // TODO: Implement
}

func (planetService) FindByID(id string) (Planet, error) {
	panic("not implemented") // TODO: Implement
}

func (planetService) FindByName(name string) (Planet, error) {
	panic("not implemented") // TODO: Implement
}

func (planetService) Add(planet Planet) (Planet, error) {
	panic("not implemented") // TODO: Implement
}

func (planetService) Delete(id string) error {
	panic("not implemented") // TODO: Implement
}

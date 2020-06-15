package planets

import (
	"starwars-hex/pkg/errs"
	"starwars-hex/pkg/logger"
)

// Service is a service interface for planets.
type Service interface {
	List() ([]Planet, error)
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

func (ps planetService) List() ([]Planet, error) {
	planets, err := ps.repo.List()
	if err != nil {
		logger.Error("planetService.List", "ps.repo.List", err)
		return planets, errs.Internal
	}

	return planets, nil
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

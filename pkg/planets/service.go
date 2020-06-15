package planets

import (
	"errors"
	"starwars-hex/pkg/errs"
	"starwars-hex/pkg/logger"
)

// Service is a service interface for planets.
type Service interface {
	List() ([]Planet, error)
	FindByName(name string) (Planet, error)
	Add(planet Planet) (Planet, error)
	Delete(name string) error
}

type planetService struct {
	repo Repository
}

// NewService creates a new planets Service.
func NewService(repo Repository) Service {
	return planetService{repo}
}

// List lists all planets.
func (ps planetService) List() ([]Planet, error) {
	planets, err := ps.repo.List()
	if err != nil {
		logger.Error("planetService.List", "ps.repo.List", err)
		return planets, errs.Unexpected
	}

	return planets, nil
}

// FindByName finds a planet by its name.
func (ps planetService) FindByName(name string) (Planet, error) {
	planet, err := ps.repo.FindByName(name)
	if err != nil {
		if errors.Is(err, errs.NoDBResults) {
			return planet, errs.NotFound
		}
		logger.Error("planetService.FindByName", "ps.repo.FindByName", err, name)
		return planet, errs.Unexpected
	}

	return planet, nil
}

// Add adds a new planet.
func (planetService) Add(planet Planet) (Planet, error) {
	panic("not implemented") // TODO: Implement
}

// Delete deletes a planet.
func (planetService) Delete(id string) error {
	panic("not implemented") // TODO: Implement
}

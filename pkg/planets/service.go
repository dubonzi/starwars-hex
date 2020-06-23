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
	repo  Repository
	swapi SwapiClient
}

// NewService creates a new planets Service.
func NewService(repo Repository, swapi SwapiClient) Service {
	return planetService{repo, swapi}
}

// List lists all planets.
func (ps planetService) List() ([]Planet, error) {
	planets, err := ps.repo.List()
	if err != nil {
		logger.Error("planetService.List", "ps.repo.List", err)
		return planets, errs.Unexpected
	}

	return planets, errs.BadRequest
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
func (ps planetService) Add(planet Planet) (Planet, error) {
	if planet.Name == "" {
		return Planet{}, errs.EmptyName
	}
	if planet.Climate == "" {
		return Planet{}, errs.EmptyClimate
	}
	if planet.Terrain == "" {
		return Planet{}, errs.EmptyTerrain
	}

	exists, err := ps.repo.Exists(planet.Name)
	if err != nil {
		logger.Error("planetService.Add", "ps.repo.Exists", err, planet.Name)
		return Planet{}, errs.Unexpected
	}
	if exists {
		return Planet{}, errs.DuplicatedPlanet
	}

	appearances, err := ps.swapi.GetFilmAppearances(planet.Name)
	if err != nil {
		logger.Error("planetService.Add", "ps.swapi.GetFilmAppearances", err, planet.Name)
		return Planet{}, errs.Unexpected
	}

	planet.FilmAppearances = appearances

	err = ps.repo.Insert(planet)
	if err != nil {
		logger.Error("planetService.Add", "ps.repo.Insert", err)
		return Planet{}, errs.Unexpected
	}

	planet, err = ps.repo.FindByName(planet.Name)
	if err != nil {
		logger.Error("planetService.Add", "ps.repo.FindByName", err, planet.Name)
		return Planet{}, errs.Unexpected
	}

	return planet, nil
}

// Delete deletes a planet.
func (ps planetService) Delete(name string) error {
	err := ps.repo.Delete(name)
	if err != nil {
		if errors.Is(err, errs.NoDBResults) {
			return errs.NotFound
		}
		logger.Error("planetService.Delete", "ps.repo.Delete", err, name)
		return errs.Unexpected
	}
	return nil
}

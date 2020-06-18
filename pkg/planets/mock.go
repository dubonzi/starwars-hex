package planets

import (
	"starwars-hex/pkg/errs"
	"starwars-hex/pkg/test"
)

// RepositoryMock is a mock for the Planet repository.
type RepositoryMock struct {
	planets []Planet

	failer test.Failer
}

// SwapiClientMock is a mock for the swapi api client.
type SwapiClientMock struct {
	appearances int

	failer test.Failer
}

// GetFilmAppearances mocks the GetFilmAppearances method.
func (s SwapiClientMock) GetFilmAppearances(name string) (int, error) {
	if s.failer != nil {
		if err := s.failer.Fails(); err != nil {
			return 0, err
		}
	}
	return s.appearances, nil
}

// List lists all planets.
func (m RepositoryMock) List() ([]Planet, error) {
	if m.failer != nil {
		if err := m.failer.Fails(); err != nil {
			return nil, err
		}
	}
	return m.planets, nil
}

// Exists checks if a planet exists.
func (m RepositoryMock) Exists(name string) (bool, error) {
	if m.failer != nil {
		if err := m.failer.Fails(); err != nil {
			return false, err
		}
	}
	for _, p := range m.planets {
		if p.Name == name {
			return true, nil
		}
	}
	return false, nil
}

// Insert inserts a new planet.
func (m *RepositoryMock) Insert(planet Planet) error {
	if m.failer != nil {
		if err := m.failer.Fails(); err != nil {
			return err
		}
	}
	m.planets = append(m.planets, planet)
	return nil
}

// FindByName finds a planet by its name (exact).
func (m RepositoryMock) FindByName(name string) (Planet, error) {
	if m.failer != nil {
		if err := m.failer.Fails(); err != nil {
			return Planet{}, err
		}
	}
	for _, p := range m.planets {
		if p.Name == name {
			return p, nil
		}
	}
	return Planet{}, errs.NoDBResults
}

// Delete deletes a planet.
func (m *RepositoryMock) Delete(name string) error {
	if m.failer != nil {
		if err := m.failer.Fails(); err != nil {
			return err
		}
	}
	for i, p := range m.planets {
		if p.Name == name {
			m.planets = append(m.planets[:i], m.planets[i+1:]...)
			return nil
		}
	}
	return errs.NoDBResults
}

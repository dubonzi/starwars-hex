package planets

import "starwars-hex/pkg/errs"

// RepositoryMock is a mock for the Planet repository.
type RepositoryMock struct {
	planets []Planet
}

// SwapiClientMock is a mock for the swapi api client.
type SwapiClientMock struct{}

// GetFilmAppearances mocks the GetFilmAppearances method.
func (SwapiClientMock) GetFilmAppearances(name string) (int, error) {
	return 1, nil
}

// List lists all planets.
func (m RepositoryMock) List() ([]Planet, error) {
	return m.planets, nil
}

// Exists checks if a planet exists.
func (m RepositoryMock) Exists(name string) (bool, error) {
	for _, p := range m.planets {
		if p.Name == name {
			return true, nil
		}
	}
	return false, nil
}

// Insert inserts a new planet.
func (m *RepositoryMock) Insert(planet Planet) error {
	m.planets = append(m.planets, planet)
	return nil
}

// FindByName finds a planet by its name (exact).
func (m RepositoryMock) FindByName(name string) (Planet, error) {
	for _, p := range m.planets {
		if p.Name == name {
			return p, nil
		}
	}
	return Planet{}, errs.NoDBResults
}

// Delete deletes a planet.
func (m *RepositoryMock) Delete(name string) error {
	for i, p := range m.planets {
		if p.Name == name {
			m.planets = append(m.planets[:i], m.planets[i+1:]...)
			return nil
		}
	}
	return errs.NoDBResults
}

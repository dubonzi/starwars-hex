package planets

import "starwars-hex/pkg/errs"

// RepositoryMock is a mock for the Planet repository.
type RepositoryMock struct {
	planets []Planet
}

// List lists all planets.
func (m RepositoryMock) List() ([]Planet, error) {
	return m.planets, nil
}

// Exists checks if a planet exists.
func (RepositoryMock) Exists(name string) (bool, error) {
	panic("not implemented") // TODO: Implement
}

// Insert inserts a new planet.
func (RepositoryMock) Insert(planet Planet) error {
	panic("not implemented") // TODO: Implement
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
func (RepositoryMock) Delete(name string) error {
	panic("not implemented") // TODO: Implement
}

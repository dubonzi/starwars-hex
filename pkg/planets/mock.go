package planets

// RepositoryMock is a mock for the Planet repository.
type RepositoryMock struct {
}

// List lists all planets.
func (RepositoryMock) List() ([]Planet, error) {
	panic("not implemented") // TODO: Implement
}

// Exists checks if a planet exists.
func (RepositoryMock) Exists(name string) (bool, error) {
	panic("not implemented") // TODO: Implement
}

// Insert inserts a new planet.
func (RepositoryMock) Insert(planet Planet) error {
	panic("not implemented") // TODO: Implement
}

// FindByID finds a planet by its ID.
func (RepositoryMock) FindByID(id string) (Planet, error) {
	panic("not implemented") // TODO: Implement
}

// FindByName finds a planet by its name (exact).
func (RepositoryMock) FindByName(name string) (Planet, error) {
	panic("not implemented") // TODO: Implement
}

// Delete deletes a planet.
func (RepositoryMock) Delete(id string) error {
	panic("not implemented") // TODO: Implement
}

package planets

// SwapiClient is the swapi api client interface.
type SwapiClient interface {
	GetFilmAppearances(string) (int, error)
}

type swapiClientImpl struct{}

// GetFilmAppearances returns the number of films a planet with `name` appeared on.
// If multiple planets are found, the first planet from the list will be chosen.
// Returns 0 if no planets are found.
func (swapiClientImpl) GetFilmAppearances(name string) (int, error) {
	panic("Not implemented")
}

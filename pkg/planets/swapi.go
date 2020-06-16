package planets

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"starwars-hex/pkg/conf"
)

// SwapiClient is the swapi api client interface.
type SwapiClient interface {
	GetFilmAppearances(string) (int, error)
}

type swapiClientImpl struct{}

// GetFilmAppearances returns the number of films a planet with `name` appeared on.
// If multiple planets are found, the first planet from the list will be chosen.
// Returns 0 if no planets are found.
func (swapiClientImpl) GetFilmAppearances(name string) (int, error) {
	client := http.Client{}
	name = url.QueryEscape(name)
	resp, err := client.Get(fmt.Sprintf("%s/planets/?search=%s", conf.SwapiURL(), name))
	if err != nil {
		return 0, err
	}

	// Using a anonymous struct since we only need information about Films from the response.
	var search struct {
		Results []struct {
			Films []string `json:"films"`
		} `json:"results"`
	}

	jsDec := json.NewDecoder(resp.Body)
	err = jsDec.Decode(&search)
	if err != nil {
		return 0, err
	}

	if len(search.Results) > 0 {
		// Using the first result because SWAPI's search uses "case-insensitive partial matches on the set of search fields".
		return len(search.Results[0].Films), nil
	}
	return 0, nil
}

package planets

import (
	"errors"
	"starwars-hex/pkg/errs"
	"testing"
)

func TestPlanetList(t *testing.T) {
	expected := []Planet{
		{
			Name:    "Naboo",
			Climate: "temperate",
			Terrain: "grassy hills, swamps, forests, mountains",
		}, {
			Name:    "Hoth",
			Climate: "frozen",
			Terrain: "tundra, ice caves, mountain ranges",
		},
	}
	svc := NewService(&RepositoryMock{planets: expected}, nil)
	planets, err := svc.List()
	if err != nil {
		t.Fatal("expected err to be nil but got: ", err)
	}
	if len(planets) != len(expected) {
		t.Errorf("expected list to have %d items, but got %d", len(expected), len(planets))
	}
}
func TestPlanetFindByName(t *testing.T) {
	data := []Planet{
		{
			Name:    "Naboo",
			Climate: "temperate",
			Terrain: "grassy hills, swamps, forests, mountains",
		}, {
			Name:    "Hoth",
			Climate: "frozen",
			Terrain: "tundra, ice caves, mountain ranges",
		},
	}

	expected := "Hoth"
	svc := NewService(&RepositoryMock{data}, nil)
	planet, err := svc.FindByName(expected)
	if err != nil {
		t.Fatal("expected err to be nil but got: ", err)
	}
	if planet.Name != expected {
		t.Errorf("expected planet name to be %s but got %s", expected, planet.Name)
	}
}

func TestPlanetNotFound(t *testing.T) {
	data := []Planet{
		{
			Name:    "Naboo",
			Climate: "temperate",
			Terrain: "grassy hills, swamps, forests, mountains",
		},
	}

	expected := errs.NotFound
	svc := NewService(&RepositoryMock{data}, nil)
	_, err := svc.FindByName("Hoth")
	if !errors.Is(err, expected) {
		t.Errorf("expected err to be '%v' but got %v", expected, err)
	}
}

func TestPlanetAddDuplicated(t *testing.T) {
	existing := []Planet{
		{
			Name:    "Naboo",
			Climate: "temperate",
			Terrain: "grassy hills, swamps, forests, mountains",
		},
	}

	svc := NewService(&RepositoryMock{existing}, nil)
	_, err := svc.Add(Planet{
		Name:    "Naboo",
		Climate: "temperate",
		Terrain: "grassy hills, swamps, forests, mountains",
	})
	if !errors.Is(err, errs.DuplicatedPlanet) {
		t.Errorf("expected err to be  '%v' but got '%v'", errs.DuplicatedPlanet, err)
	}

}

func TestPlanetAddInvalid(t *testing.T) {
	svc := NewService(&RepositoryMock{}, nil)

	_, err := svc.Add(Planet{})
	if !errors.Is(err, errs.EmptyName) {
		t.Errorf("expected err to be '%v' but got '%v'", errs.EmptyName, err)
	}

	_, err = svc.Add(Planet{
		Name: "Hoth",
	})
	if !errors.Is(err, errs.EmptyClimate) {
		t.Errorf("expected err to be '%v' but got '%v'", errs.EmptyClimate, err)
	}

	_, err = svc.Add(Planet{
		Name:    "Hoth",
		Climate: "frozen",
	})
	if !errors.Is(err, errs.EmptyTerrain) {
		t.Errorf("expected err to be '%v' but got '%v'", errs.EmptyTerrain, err)
	}
}

func TestPlanetAdd(t *testing.T) {
	svc := NewService(&RepositoryMock{}, SwapiClientMock{})

	expected := Planet{
		Name:    "Naboo",
		Climate: "temperate",
		Terrain: "grassy hills, swamps, forests, mountains",
	}

	planet, err := svc.Add(expected)
	if err != nil {
		t.Fatal("expected err to be nil but got: ", err)
	}

	if planet.Name != expected.Name {
		t.Errorf("expected planet name to be '%s' but got '%s'", expected.Name, planet.Name)
	}

	if planet.Climate != expected.Climate {
		t.Errorf("expected planet climate to be '%s' but got '%s'", expected.Climate, planet.Climate)
	}

	if planet.Terrain != expected.Terrain {
		t.Errorf("expected planet terrain to be '%s' but got '%s'", expected.Terrain, planet.Terrain)
	}
}

func TestPlanetDelete(t *testing.T) {
	data := []Planet{
		{
			Name:    "Naboo",
			Climate: "temperate",
			Terrain: "grassy hills, swamps, forests, mountains",
		}, {
			Name:    "Hoth",
			Climate: "frozen",
			Terrain: "tundra, ice caves, mountain ranges",
		},
		{
			Name:    "Tatooine",
			Climate: "arid",
			Terrain: "desert",
		},
		{
			Name:    "Alderaan",
			Climate: "temperate",
			Terrain: "grasslands, mountains",
		},
	}
	mock := &RepositoryMock{planets: data}
	svc := NewService(mock, nil)

	toDelete := "Hoth"
	err := svc.Delete(toDelete)
	if err != nil {
		t.Fatal("expected err to be nil but got ", err)
	}

	if len(data) == len(mock.planets) {
		t.Errorf("expected list to have fewer than %d items", len(data))
	}

	if _, err = svc.FindByName(toDelete); err == nil {
		t.Errorf("expected planet '%s' to be deleted, but it was found in the list", toDelete)
	}

}

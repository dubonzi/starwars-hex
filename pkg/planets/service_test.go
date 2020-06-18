package planets

import (
	"errors"
	"starwars-hex/pkg/errs"
	"starwars-hex/pkg/test"
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

func TestPlanetFailingList(t *testing.T) {
	expected := []Planet{
		{
			Name:    "Naboo",
			Climate: "temperate",
			Terrain: "grassy hills, swamps, forests, mountains",
		},
	}
	svc := NewService(&RepositoryMock{
		planets: expected,
		failer:  test.RandomFail{FailRate: 0.1},
	}, nil)

	if _, err := svc.List(); err != nil {
		_, ok := err.(*errs.HTTPError)
		if !ok {
			t.Errorf("expected err to be of type *errs.HTTPError but got %T", err)
		}
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
	svc := NewService(&RepositoryMock{planets: data}, nil)
	planet, err := svc.FindByName(expected)
	if err != nil {
		t.Fatal("expected err to be nil but got: ", err)
	}
	if planet.Name != expected {
		t.Errorf("expected planet name to be %s but got %s", expected, planet.Name)
	}
}

func TestPlanetFailingFindByName(t *testing.T) {
	data := []Planet{
		{
			Name:    "Naboo",
			Climate: "temperate",
			Terrain: "grassy hills, swamps, forests, mountains",
		},
	}

	svc := NewService(&RepositoryMock{
		planets: data,
		failer:  test.RandomFail{FailRate: 0.1},
	}, nil)

	if _, err := svc.FindByName("Naboo"); err != nil {
		_, ok := err.(*errs.HTTPError)
		if !ok {
			t.Errorf("expected err to be of type *errs.HTTPError but got %T", err)
		}
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
	svc := NewService(&RepositoryMock{planets: data}, nil)
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

	svc := NewService(&RepositoryMock{planets: existing}, nil)
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
	expectedAppearances := 5
	expected := Planet{
		Name:    "Naboo",
		Climate: "temperate",
		Terrain: "grassy hills, swamps, forests, mountains",
	}

	svc := NewService(&RepositoryMock{}, SwapiClientMock{appearances: expectedAppearances})

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

	if planet.FilmAppearances != expectedAppearances {
		t.Errorf("expected appearances to be %d but got %d", expectedAppearances, planet.FilmAppearances)
	}
}

func TestPlanetFailingAdd(t *testing.T) {
	svc := NewService(&RepositoryMock{
		failer: test.RandomFail{FailRate: 0.1},
	}, SwapiClientMock{})

	planet := Planet{
		Name:    "Naboo",
		Climate: "temperate",
		Terrain: "grassy hills, swamps, forests, mountains",
	}

	_, err := svc.Add(planet)
	if err != nil {
		_, ok := err.(*errs.HTTPError)
		if !ok {
			t.Errorf("expected err to be of type *errs.HTTPError but got %T", err)
		}
	}
}

func TestSwapiFailingGetAppearances(t *testing.T) {
	svc := NewService(&RepositoryMock{}, SwapiClientMock{
		failer: test.RandomFail{FailRate: 0.1},
	})

	planet := Planet{
		Name:    "Naboo",
		Climate: "temperate",
		Terrain: "grassy hills, swamps, forests, mountains",
	}

	_, err := svc.Add(planet)
	if err != nil {
		_, ok := err.(*errs.HTTPError)
		if !ok {
			t.Errorf("expected err to be of type *errs.HTTPError but got %T", err)
		}
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

func TestPlanetDeleteNotFound(t *testing.T) {
	data := []Planet{
		{
			Name:    "Naboo",
			Climate: "temperate",
			Terrain: "grassy hills, swamps, forests, mountains",
		},
	}
	mock := &RepositoryMock{planets: data}
	svc := NewService(mock, nil)

	toDelete := "Hoth"
	err := svc.Delete(toDelete)
	if !errors.Is(err, errs.NotFound) {
		t.Errorf("expected err to be *errs.NotFound but got %T", err)
	}
}

func TestPlanetFailingDelete(t *testing.T) {
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
	mock := &RepositoryMock{
		planets: data,
		failer:  test.RandomFail{FailRate: 0.1},
	}
	svc := NewService(mock, nil)

	if err := svc.Delete("Naboo"); err != nil {
		_, ok := err.(*errs.HTTPError)
		if !ok {
			t.Errorf("expected err to be of type *errs.HTTPError but got %T", err)
		}
	}
}

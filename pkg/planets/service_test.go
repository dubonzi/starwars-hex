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
	mock := RepositoryMock{
		planets: expected,
	}
	svc := NewService(mock)
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
	mock := RepositoryMock{data}
	svc := NewService(mock)
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
	mock := RepositoryMock{data}
	svc := NewService(mock)
	_, err := svc.FindByName("Hoth")
	if !errors.Is(err, expected) {
		t.Errorf("expected err to be '%v' but got %v", expected, err)
	}
}

func TestPlanetAdd(t *testing.T) {
	mock := RepositoryMock{}
	svc := NewService(mock)
	svc.Add(Planet{})
}
func TestPlanetDelete(t *testing.T) {
	mock := RepositoryMock{}
	svc := NewService(mock)
	svc.Delete("")
}

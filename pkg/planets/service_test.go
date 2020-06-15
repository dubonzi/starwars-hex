package planets

import "testing"

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
		t.Error("expected err to be nil but got: ", err)
	}
	if len(planets) != len(expected) {
		t.Errorf("expected list to have %d items, but got %d", len(expected), len(planets))
	}
}
func TestPlanetAdd(t *testing.T) {
	mock := RepositoryMock{}
	svc := NewService(mock)
	svc.Add(Planet{})
}
func TestPlanetFindByName(t *testing.T) {
	mock := RepositoryMock{}
	svc := NewService(mock)
	svc.FindByName("")
}
func TestPlanetDelete(t *testing.T) {
	mock := RepositoryMock{}
	svc := NewService(mock)
	svc.Delete("")
}

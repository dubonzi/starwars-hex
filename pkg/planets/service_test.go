package planets

import "testing"

func TestPlanetList(t *testing.T) {
	mock := RepositoryMock{}
	svc := NewService(mock)
	svc.List()
}
func TestPlanetAdd(t *testing.T) {
	mock := RepositoryMock{}
	svc := NewService(mock)
	svc.Add(Planet{})
}
func TestPlanetFindByID(t *testing.T) {
	mock := RepositoryMock{}
	svc := NewService(mock)
	svc.FindByID("")
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

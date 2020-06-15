package errs

import (
	"net/http"
)

// HTTPError is an error message intended to be sent to the api clients.
type HTTPError struct {
	Status  int
	Code    string
	Message string
}

// Error implements the Error interface.
func (e *HTTPError) Error() string {
	return e.Message
}

var (
	//Internal represents an internal server error.
	Internal = &HTTPError{Status: http.StatusInternalServerError, Code: "internal_server_error", Message: "An unexpected internal server error occurred."}
	//BadRequest represents a generic bad request error.
	BadRequest = &HTTPError{Status: http.StatusBadRequest, Code: "bad_request", Message: "The information sent was in an invalid format."}
	//NotFound represents a not found error.
	NotFound = &HTTPError{Status: http.StatusNotFound, Code: "not_found", Message: "The requested resource was not found."}

	//InvalidID represents an invalid ID error.
	InvalidID = &HTTPError{Status: http.StatusBadRequest, Code: "invalid.id", Message: "The given ID is invalid."}

	//EmptyName represents an empty name error.
	EmptyName = &HTTPError{Status: http.StatusBadRequest, Code: "name.empty", Message: "The name cannot be empty."}
	//EmptyClimate represents an empty climate error.
	EmptyClimate = &HTTPError{Status: http.StatusBadRequest, Code: "climate.empty", Message: "The climate cannot be empty."}
	//EmptyTerrain represents an empty terrain error.
	EmptyTerrain = &HTTPError{Status: http.StatusBadRequest, Code: "terrain.empty", Message: "The terrain cannot be empty."}

	//DuplicatedPlanet represents an error for duplicated planets.
	DuplicatedPlanet = &HTTPError{Status: http.StatusBadRequest, Code: "planet.duplicated", Message: "A planet with the given name already exists."}
)

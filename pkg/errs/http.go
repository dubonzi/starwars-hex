package errs

import (
	"net/http"

	"github.com/labstack/echo/v4"
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

// NewHTTPErrorFromEcho creates a new HTTPError from a *echo.HTTPError.
func NewHTTPErrorFromEcho(err *echo.HTTPError) *HTTPError {
	echoMsg, ok := err.Message.(string)
	if !ok {
		return Unexpected
	}
	return &HTTPError{
		Status:  err.Code,
		Message: echoMsg,
	}
}

var (
	//Unexpected represents an internal server error.
	Unexpected = &HTTPError{Status: http.StatusInternalServerError, Code: "internal_server_error", Message: "An unexpected internal server error occurred."}
	//BadRequest represents a generic bad request error.
	BadRequest = &HTTPError{Status: http.StatusBadRequest, Code: "bad_request", Message: "The information sent was in an invalid format."}
	//NotFound represents a not found error.
	NotFound = &HTTPError{Status: http.StatusNotFound, Code: "not_found", Message: "The requested resource was not found."}

	//EmptyName represents an empty name error.
	EmptyName = &HTTPError{Status: http.StatusBadRequest, Code: "name.empty", Message: "The name cannot be empty."}
	//EmptyClimate represents an empty climate error.
	EmptyClimate = &HTTPError{Status: http.StatusBadRequest, Code: "climate.empty", Message: "The climate cannot be empty."}
	//EmptyTerrain represents an empty terrain error.
	EmptyTerrain = &HTTPError{Status: http.StatusBadRequest, Code: "terrain.empty", Message: "The terrain cannot be empty."}

	//DuplicatedPlanet represents an error for duplicated planets.
	DuplicatedPlanet = &HTTPError{Status: http.StatusBadRequest, Code: "planet.duplicated", Message: "A planet with the given name already exists."}
)

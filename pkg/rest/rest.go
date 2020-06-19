package rest

import (
	"encoding/json"
	"net/http"
	"starwars-hex/pkg/errs"
)

// JSON encodes data as JSON and writes it to w.
func JSON(w http.ResponseWriter, status int, data interface{}) error {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(status)
	enc := json.NewEncoder(w)
	if err := enc.Encode(data); err != nil {
		return errs.Unexpected
	}
	return nil
}

// SendError encodes err as JSON and writes it to w.
//	If err is not of type *errs.HTTPError, errs.Unexpected will be used.
func SendError(w http.ResponseWriter, err error) error {
	herr, ok := err.(*errs.HTTPError)
	if !ok {
		herr = errs.Unexpected
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(herr.Status)
	enc := json.NewEncoder(w)
	return enc.Encode(herr)
}

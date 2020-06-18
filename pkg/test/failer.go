package test

import (
	"errors"
	"math/rand"
)

// Failer can be implemented to cause an error randomly or with a fixed rate.
type Failer interface {
	Fails() error
}

// RandomFail randomly fails based on a fail rate.
type RandomFail struct {
	FailRate float32
}

// Fails implement the test.Failer interface.
func (r RandomFail) Fails() error {
	if rand.Float32() < r.FailRate {
		return errors.New("failed randomly")
	}
	return nil
}

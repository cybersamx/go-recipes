package services

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type stateNotFoundError struct {
	abbreviation string
}

func newStateNotFoundError(abbreviation string) stateNotFoundError {
	return stateNotFoundError{
		abbreviation: abbreviation,
	}
}

// Implements error
func (snfe stateNotFoundError) Error() string {
	return fmt.Sprintf("can't find state %s", snfe.abbreviation)
}

// Implements transport.StatusCoder to tell go-kit to use http status code specified
// in the StatusCode() method instead of the default status code 500.
func (snfe stateNotFoundError) StatusCode() int {
	return http.StatusNotFound
}

// Implements json.Marshaler to tell go-kit to return the content as json according to the
// MarshalJSON() method.
func (snfe stateNotFoundError) MarshalJSON() ([]byte, error) {
	content := struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	}{
		Code:    snfe.StatusCode(),
		Message: snfe.Error(),
	}

	return json.Marshal(content)
}

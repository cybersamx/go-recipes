package main

import (
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

func TestStringValidation(t *testing.T) {
	validate = validator.New()

	// Returns a map[string]string type that implements the error interface.
	errs := validate.Var("sam@example.com", "email")
	assert.Nil(t, errs)

	errs = validate.Var(20, "gte=12,lte=120")
	assert.Nil(t, errs)
}

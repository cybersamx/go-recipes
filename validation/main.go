package main

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

var (
	validate *validator.Validate
)

func main() {
	validate = validator.New()

	// We deliberately let the validations fail.

	// validate.Var returns a map[string]string type that implements the error interface.
	errs := validate.Var("sam@example", "email")
	if errs != nil {
		fmt.Println(errs)
	}

	errs = validate.Var(5, "gte=12,lte=120")
	if errs != nil {
		fmt.Println(errs)
	}
}

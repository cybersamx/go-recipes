// CREDIT: Based loosely on the sample code found:
// https://github.com/go-playground/validator/blob/master/_examples/simple/main.go

package main

import (
	"fmt"
	"log"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	ent "github.com/go-playground/validator/v10/translations/en"
)

type Person struct {
	Name  string            `validate:"required"`
	Email string            `validate:"email"`
	Age   int               `validate:"gte=12,lte=120"`
	Tags  map[string]string `validate:"dive,required"`
}

func checkError(err error) {
	verrs, ok := err.(validator.ValidationErrors)
	if !ok {
		log.Fatal("failed to cast validator.ValidationError")
	}

	// validate.Var and validateStruct return a validator.ValidationErrors type, which
	// an type alias for a slice-custom-type that also implements the error interface. So we can
	// validator.ValidationErrors as a slice-custom-type and error.
	for _, verr := range verrs {
		fmt.Printf("Namespace=%s StructNamespace=%s Field=%s StructField=%s\n", verr.Namespace(), verr.StructNamespace(), verr.Field(), verr.StructField())
		fmt.Printf("Tag=%s Kind=%s Type=%s Value=%s Param=%s\n\n", verr.Tag(), verr.Kind(), verr.Type(), verr.Value(), verr.Param())
	}
}

func validateVar(validate *validator.Validate) {
	email := "sam@example"
	age := 5

	// We deliberately let the validations fail.

	// validate.Var returns a map[string]string type that implements the error interface.
	checkError(validate.Var(email, "email"))

	checkError(validate.Var(age, "gte=12,lte=120"))
}

func validateStruct(validate *validator.Validate) {
	person := Person{
		Name:  "",
		Email: "sam@example",
		Age:   5,
		Tags: map[string]string{
			"Job": "",
		},
	}

	if err := validate.Struct(person); err != nil {
		// We may pass a nil or invalid struct value to validate.Struct and cause it to return
		// an error of type InvalidValidationError. Consider checking for InvalidValidationError.
		if _, ok := err.(*validator.InvalidValidationError); ok {
			fmt.Printf("Invalid or nil struct value: %v", err)
			return
		}

		checkError(err)
	}
}

func validateWithErrTranslator(validate *validator.Validate) {
	// Set up translator
	locale := "en"
	en := en.New()
	uni := ut.New(en, en)
	trans, ok := uni.GetTranslator(locale)
	if !ok {
		log.Fatalf("can't find %s locale", locale)
	}
	ent.RegisterDefaultTranslations(validate, trans)

	person := Person{}

	if err := validate.Struct(person); err != nil {
		verrs, ok := err.(validator.ValidationErrors)
		if !ok {
			log.Fatal("failed to cast validator.ValidationError")
		}

		// Unlike checkError, we are using the translator to print out a more human
		// readable error message in the assigned locale/language.
		for _, verr := range verrs {
			fmt.Println(verr.Translate(trans))
		}
	}
}

func main() {
	validate := validator.New()

	validateVar(validate)
	validateStruct(validate)
	validateWithErrTranslator(validate)
}

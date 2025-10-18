package validator

import (
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

func ValidateStruct(v any) map[string]string {
	errors := make(map[string]string)
	validate := validator.New()
	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		name := strings.SplitN(field.Tag.Get("json"), ",", 2)[0]

		if name == "-" {
			return ""
		}
		return name
	})

	err := validate.Struct(v)
	if err == nil {
		return nil
	}

	errs := err.(validator.ValidationErrors)
	for _, e := range errs {
		errors[e.Field()] = e.Tag()
	}

	return errors
}

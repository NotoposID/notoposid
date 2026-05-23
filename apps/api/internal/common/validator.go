package common

import "github.com/go-playground/validator/v10"

var Validate = validator.New()

type ErrorResponseItem struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func ValidateStruct(s interface{}) []*ErrorResponseItem {
	var errors []*ErrorResponseItem
	err := Validate.Struct(s)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponseItem
			element.Field = err.StructNamespace()
			switch err.Tag() {
			case "required":
				element.Message = "Kolom " + err.Field() + " wajib diisi"
			case "email":
				element.Message = "Kolom " + err.Field() + " harus berupa format email yang valid"
			case "gt":
				element.Message = "Kolom " + err.Field() + " harus lebih besar dari " + err.Param()
			case "min":
				element.Message = "Kolom " + err.Field() + " minimal " + err.Param() + " karakter"
			default:
				element.Message = "Kolom " + err.Field() + " tidak valid (" + err.Tag() + ")"
			}
			errors = append(errors, &element)
		}
	}
	return errors
}

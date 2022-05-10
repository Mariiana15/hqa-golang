package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc

type MetaData interface{}

type Car struct {
	Id          string `json:"id"`
	Brand       string `json:"brand" validate:"required,alphanum"`
	Model       string `json:"model" validate:"required,alphanum"`
	Horse_power uint32 `json:"horse_power" validate:"required,gte=0,lte=10000"`
}

var validate *validator.Validate

func (car *Car) ToJson() ([]byte, error) {
	return json.Marshal(car)
}

func (car *Car) ValidateStructure() (string, error) {
	validate = validator.New()
	var msg string
	err := validate.Struct(car)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return err.Error(), err
		}
		return fmt.Sprintf("The field %s is mal format", err.(validator.ValidationErrors)[0].StructField()), err
	}
	return msg, err
}

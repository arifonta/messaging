package customvalidator

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
)

type CustomValidator struct {
	Validator *validator.Validate
}

func NewCustomValidator() *CustomValidator {
	return &CustomValidator{Validator: validator.New()}
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.Validator.Struct(i)
}

func CustomErrorMessage(err error) error {

	var reportMessage interface{}
	// fmt.Println("error validation")
	if castedObject, ok := err.(validator.ValidationErrors); ok {
		var message []string
		for _, err := range castedObject {
			switch err.Tag() {
			case "required":
				message = append(message, fmt.Sprintf("%s is required",
					err.Field()))
			case "email":
				message = append(message, fmt.Sprintf("%s is not valid email",
					err.Field()))
			case "gte":
				message = append(message, fmt.Sprintf("%s value must be greater than %s",
					err.Field(), err.Param()))
			case "lte":
				message = append(message, fmt.Sprintf("%s value must be lower than %s",
					err.Field(), err.Param()))
			}
		}
		// report.Message = message
		reportMessage = message
	}

	return errors.New(
		fmt.Sprint(reportMessage),
	)
}

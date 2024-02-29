package utils

import (
	"fmt"
	"strings"

	"github.com/MuhammedYahiya/Ecom-api/pkg/domain"
	"github.com/go-playground/validator/v10"
)

func ValidateUser(user domain.User) error {
	validate := validator.New()

	err := validate.Struct(user)
	if err != nil {
		validateErrors := err.(validator.ValidationErrors)
		errorMessage := make([]string, len(validateErrors))

		for i, validateErr := range validateErrors {
			fieldName := validateErr.Field()
			switch fieldName {
			case "Username":
				errorMessage[i] = "Invalid Username, Minimum 8 letters Maximum 24 letters"
			case "Email":
				errorMessage[i] = "Invalid Email"
			case "Phone":
				errorMessage[i] = "Invalid phone number"
			case "Password":
				errorMessage[i] = "Invalid password, Minimum 8 letters or Maximum 16 letters required"
			default:
				errorMessage[i] = "Validation failed"
			}
		}
		return fmt.Errorf(strings.Join(errorMessage, ", "))

	}

	return nil
}

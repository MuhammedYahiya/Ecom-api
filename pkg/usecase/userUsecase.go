package usecase

import (
	"github.com/MuhammedYahiya/Ecom-api/pkg/domain"
	"github.com/MuhammedYahiya/Ecom-api/pkg/repository"
	"github.com/MuhammedYahiya/Ecom-api/pkg/utils"
)

func CreateUser(userData *domain.User) error {
	validateErr := utils.ValidateUser(*userData)
	if validateErr != nil {
		return validateErr
	}
	err := repository.CreateUser(userData)
	if err != nil {
		return err
	}
	return nil
}

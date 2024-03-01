package usecase

import (
	"errors"

	"github.com/MuhammedYahiya/Ecom-api/pkg/domain"
	"github.com/MuhammedYahiya/Ecom-api/pkg/repository"
	"github.com/MuhammedYahiya/Ecom-api/pkg/utils"
)

func CreateUser(userData *domain.User) error {
	validateErr := utils.ValidateUser(*userData)
	if validateErr != nil {
		return validateErr
	}
	res, err := repository.FindUserByEmail(userData)
	if err != nil && res == nil {
		pass, error := utils.HashPassword(userData.Password)
		if error != nil {
			return errors.New("failed to hash")
		}
		userData.Password = pass
		err := repository.CreateUser(userData)
		return err
	}
	return errors.New("user  with the same mail id  already exist")
}

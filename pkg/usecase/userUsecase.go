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
	res, err := repository.FindUserByEmail(userData.Email)
	if err != nil && res == nil {
		otp, otpError := utils.Otpgeneration(userData.Email)
		if otpError != nil {
			return otpError
		}
		userData.Otp = otp
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

func RegistrationValidate(userData *domain.User) error {
	enteredOtp := userData.Otp
	res, err := repository.FindUserByEmail(userData.Email)
	if err != nil {
		return errors.New("you should register first")
	}
	if userData.Email == res.Email && enteredOtp == res.Otp {
		return nil
	}
	return errors.New("invalid Otp")

}

func VerifyUser(userData *domain.User) error {
	err := repository.UpdateUserStatus(userData.Email)
	if err != nil {
		return err
	}
	return nil
}

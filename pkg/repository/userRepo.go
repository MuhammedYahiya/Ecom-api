package repository

import (
	"github.com/MuhammedYahiya/Ecom-api/pkg/db"
	"github.com/MuhammedYahiya/Ecom-api/pkg/domain"
	"gorm.io/gorm"
)

func CreateUser(userData *domain.User) error {
	err := db.DB.Create(&userData)
	if err != nil {
		return err.Error
	}
	return nil
}

func FindUserByEmail(userData *domain.User) (*domain.User, error) {
	user := &domain.User{}

	result := db.DB.Where("email=?", userData.Email).First(userData)

	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	return user, nil
}

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

func FindUserByEmail(email string) (*domain.User, error) {
	user := &domain.User{}

	result := db.DB.Where("email=?", email).First(user)

	if result.Error != nil {
		if result.RowsAffected == 0 {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, result.Error
	}

	return user, nil
}

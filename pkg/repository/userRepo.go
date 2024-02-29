package repository

import (
	"github.com/MuhammedYahiya/Ecom-api/pkg/db"
	"github.com/MuhammedYahiya/Ecom-api/pkg/domain"
)

func CreateUser(userData *domain.User) error {
	err := db.DB.Create(&userData)
	if err != nil {
		return err.Error
	}
	return nil
}

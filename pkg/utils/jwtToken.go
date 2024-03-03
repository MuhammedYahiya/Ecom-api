package utils

import (
	"os"
	"time"

	"github.com/MuhammedYahiya/Ecom-api/pkg/domain"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateJwtToken(userData domain.User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["role"] = userData.IsAdmin
	claims["email"] = userData.Email
	claims["userid"] = userData.ID
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	claims["iat"] = time.Now().Unix()

	secretKey := []byte(os.Getenv("JWT_secret_key"))
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

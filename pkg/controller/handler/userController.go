package handler

import (
	"net/http"

	"github.com/MuhammedYahiya/Ecom-api/pkg/domain"
	"github.com/MuhammedYahiya/Ecom-api/pkg/usecase"
	"github.com/gin-gonic/gin"
)

func RegisterUser(ctx *gin.Context) {
	userData := domain.User{}

	if err := ctx.ShouldBindJSON(&userData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"Message": "User data not bind",
			"Error":   err,
		})
		return
	}

	err := usecase.CreateUser(&userData)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"Message": "Register the User Failed",
			"Error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"Message": "Redirect: http://localhost:8080/user/register/validate",
		"Error":   nil,
	})
}

func ValidateUser(ctx *gin.Context) {
	userData := domain.User{}
	if err := ctx.ShouldBindJSON(&userData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"Message": "User data not bind",
			"Error":   err,
		})
		return
	}

	err := usecase.RegistrationValidate(&userData)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"Message": "User validation failed",
			"Error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"Message": "User Validation success",
		"Error":   nil,
	})

}

package controller

import (
	"github.com/MuhammedYahiya/Ecom-api/pkg/controller/handler"
	"github.com/gin-gonic/gin"
)

func InitializeRouter(r *gin.Engine) {

	userGroup := r.Group("/user")
	{
		userGroup.POST("/register", handler.RegisterUser)
	}

}

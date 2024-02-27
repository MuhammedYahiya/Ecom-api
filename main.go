package main

import (
	"net/http"

	"github.com/MuhammedYahiya/Ecom-api/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	utils.ConnectDb()

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, Gin!")
	})

	router.Run(":8080")
}

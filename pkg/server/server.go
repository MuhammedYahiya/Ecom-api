package server

import "github.com/gin-gonic/gin"

func ServerConnect() *gin.Engine {
	r := gin.Default()
	return r
}

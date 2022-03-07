package server

import "github.com/gin-gonic/gin"

func initRouter() *gin.Engine {
	router := gin.Default()

	return router
}

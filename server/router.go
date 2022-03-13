package server

import "github.com/gin-gonic/gin"

func InitRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/covid/summary", getCovidSummaryHandler)

	return router
}

package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getCovidSummaryHandler(c *gin.Context) {
	res := getCovidSummary()

	c.JSON(http.StatusOK, res)
}

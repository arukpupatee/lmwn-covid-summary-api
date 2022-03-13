package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getCovidSummaryHandler(c *gin.Context) {
	res := GetCovidSummary()

	c.JSON(http.StatusOK, res)
}

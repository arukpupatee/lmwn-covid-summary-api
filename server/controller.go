package server

import (
	"net/http"

	"github.com/arukpupatee/lmwn-covid-summary-api/services"
	"github.com/gin-gonic/gin"
)

func getCovidSummaryHandler(c *gin.Context) {
	res := services.GetCovidSummary()

	c.JSON(http.StatusOK, res)
}

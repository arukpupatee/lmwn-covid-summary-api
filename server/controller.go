package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type province map[string]int
type ageGroup struct {
	Young   int `json:"0-30"`
	Middle  int `json:"31-60"`
	Old     int `json:"61+"`
	Unknown int `json:"N/A"`
}
type Response struct {
	Province province
	AgeGroup ageGroup
}

func getCovidSummaryHandler(c *gin.Context) {
	res := Response{
		Province: map[string]int{
			"Samut Sakhon": 3613,
			"Bangkok":      2774,
		},
		AgeGroup: ageGroup{
			Young:   300,
			Middle:  150,
			Old:     250,
			Unknown: 4,
		},
	}

	c.JSON(http.StatusOK, res)
}

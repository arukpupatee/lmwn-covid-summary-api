package services

import (
	"github.com/arukpupatee/lmwn-covid-summary-api/external"
	"github.com/arukpupatee/lmwn-covid-summary-api/models"
)

const unknownAgeValue = 0
const youngUpperBoundValue = 30
const middleUpperBoundValue = 60

func GetCovidSummary() models.Summary {
	// TODO: Handle error
	covidCaseResponseBody, _ := external.FetchCovidCase()
	covidCaseData := covidCaseResponseBody.Data

	provinceData := models.Province{}
	ageGroupData := models.AgeGroup{}

	for _, data := range covidCaseData {
		province := data.ProvinceEn
		age := data.Age

		if province != "" {
			if _, ok := provinceData[province]; !ok {
				provinceData[province] = 1
			} else {
				provinceData[province]++
			}
		}

		if age == unknownAgeValue {
			ageGroupData.Unknown++
		} else if age <= youngUpperBoundValue {
			ageGroupData.Young++
		} else if age <= middleUpperBoundValue {
			ageGroupData.Middle++
		} else {
			ageGroupData.Old++
		}
	}

	summary := models.Summary{
		Province: provinceData,
		AgeGroup: ageGroupData,
	}

	return summary
}

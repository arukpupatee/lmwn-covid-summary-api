package server

import "github.com/arukpupatee/lmwn-covid-summary-api/external"

type province map[string]int
type AgeGroup struct {
	Young   int `json:"0-30"`
	Middle  int `json:"31-60"`
	Old     int `json:"61+"`
	Unknown int `json:"N/A"`
}
type Summary struct {
	Province province
	AgeGroup AgeGroup
}

const unknownAgeValue = 0
const youngUpperBoundValue = 30
const middleUpperBoundValue = 60

func GetCovidSummary() Summary {
	// TODO: Handle error
	covidCaseResponseBody, _ := external.FetchCovidCase()
	covidCaseData := covidCaseResponseBody.Data

	provinceData := province{}
	ageGroupData := AgeGroup{}

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

	summary := Summary{
		Province: provinceData,
		AgeGroup: ageGroupData,
	}

	return summary
}

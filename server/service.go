package server

type province map[string]int
type ageGroup struct {
	Young   int `json:"0-30"`
	Middle  int `json:"31-60"`
	Old     int `json:"61+"`
	Unknown int `json:"N/A"`
}
type Summary struct {
	Province province
	AgeGroup ageGroup
}

const unknownAgeValue = 0
const youngUpperBoundValue = 30
const middleUpperBoundValue = 60

func getCovidSummary() Summary {
	// TODO: Handle error
	covidCaseResponseBody, _ := fetchCovidCase()
	covidCaseData := covidCaseResponseBody.Data

	provinceData := province{}
	ageGroupData := ageGroup{}

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

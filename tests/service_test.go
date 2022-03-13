package tests

import (
	"testing"

	"github.com/arukpupatee/lmwn-covid-summary-api/models"
	"github.com/arukpupatee/lmwn-covid-summary-api/services"
)

func TestGetCovidSummary(t *testing.T) {
	summary := services.GetCovidSummary()

	expectedSummary := models.Summary{
		Province: map[string]int{
			"Amnat Charoen":            17,
			"Ang Thong":                36,
			"Bangkok":                  27,
			"Bueng Kan":                23,
			"Buriram":                  18,
			"Chachoengsao":             24,
			"Chai Nat":                 25,
			"Chaiyaphum":               28,
			"Chanthaburi":              17,
			"Chiang Mai":               22,
			"Chiang Rai":               15,
			"Chonburi":                 29,
			"Chumphon":                 25,
			"Kalasin":                  27,
			"Kamphaeng Phet":           23,
			"Kanchanaburi":             23,
			"Khon Kaen":                27,
			"Krabi":                    27,
			"Lampang":                  24,
			"Lamphun":                  25,
			"Loei":                     17,
			"Lopburi":                  19,
			"Mae Hong Son":             22,
			"Maha Sarakham":            26,
			"Mukdahan":                 28,
			"Nakhon Nayok":             19,
			"Nakhon Pathom":            31,
			"Nakhon Phanom":            24,
			"Nakhon Ratchasima":        28,
			"Nakhon Sawan":             24,
			"Nakhon Si Thammarat":      35,
			"Nan":                      20,
			"Narathiwat":               22,
			"Nong Bua Lamphu":          29,
			"Nong Khai":                27,
			"Nonthaburi":               29,
			"Pathum Thani":             30,
			"Pattani":                  27,
			"Phang Nga":                28,
			"Phatthalung":              29,
			"Phayao":                   25,
			"Phetchabun":               33,
			"Phetchaburi":              26,
			"Phichit":                  21,
			"Phitsanulok":              24,
			"Phra Nakhon Si Ayutthaya": 25,
			"Phrae":                    28,
			"Phuket":                   25,
			"Prachinburi":              19,
			"Prachuap Khiri Khan":      34,
			"Ranong":                   35,
			"Ratchaburi":               21,
			"Rayong":                   25,
			"Roi Et":                   25,
			"Sa Kaeo":                  26,
			"Sakon Nakhon":             42,
			"Samut Prakan":             31,
			"Samut Sakhon":             29,
			"Samut Songkhram":          22,
			"Saraburi":                 26,
			"Satun":                    37,
			"Sing Buri":                26,
			"Sisaket":                  27,
			"Songkhla":                 24,
			"Sukhothai":                23,
			"Suphan Buri":              28,
			"Surat Thani":              25,
			"Surin":                    24,
			"Tak":                      18,
			"Trang":                    20,
			"Trat":                     25,
			"Ubon Ratchathani":         23,
			"Udon Thani":               34,
			"Uthai Thani":              24,
			"Uttaradit":                24,
			"Yala":                     27,
			"Yasothon":                 26,
		},
		AgeGroup: models.AgeGroup{
			Young:   582,
			Middle:  607,
			Old:     769,
			Unknown: 42,
		},
	}

	for province := range expectedSummary.Province {
		if expectedSummary.Province[province] != summary.Province[province] {
			t.Fatalf("%v province is %v, expect %v", province, summary.Province[province], expectedSummary.Province[province])
		}
	}

	// TODO: Refactor by iterate over ageGroup struct field
	if expectedSummary.AgeGroup.Young != summary.AgeGroup.Young {
		t.Fatalf("Young age group is %v, expect %v", summary.AgeGroup.Young, expectedSummary.AgeGroup.Young)
	}
	if expectedSummary.AgeGroup.Middle != summary.AgeGroup.Middle {
		t.Fatalf("Middle age group is %v, expect %v", summary.AgeGroup.Middle, expectedSummary.AgeGroup.Middle)
	}
	if expectedSummary.AgeGroup.Old != summary.AgeGroup.Old {
		t.Fatalf("Old age group is %v, expect %v", summary.AgeGroup.Old, expectedSummary.AgeGroup.Old)
	}
	if expectedSummary.AgeGroup.Unknown != summary.AgeGroup.Unknown {
		t.Fatalf("Unknown age group is %v, expect %v", summary.AgeGroup.Unknown, expectedSummary.AgeGroup.Unknown)
	}
}

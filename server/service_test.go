package server

import (
	"testing"
)

func TestGetCovidSummary(t *testing.T) {
	summary := getCovidSummary()

	expectedSummary := Summary{
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

package tests

import (
	"testing"

	"github.com/arukpupatee/lmwn-covid-summary-api/external"
)

func TestFetchCovidCaseSuccess(t *testing.T) {
	_, err := external.FetchCovidCase()

	if err != nil {
		t.Fatalf("Error '%v' when fetch covid case from external data source", err)
	}
}

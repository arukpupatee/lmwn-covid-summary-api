package external

import "testing"

func TestFetchCovidCaseSuccess(t *testing.T) {
	_, err := fetchCovidCase()

	if err != nil {
		t.Fatalf("Error '%v' when fetch covid case from external data source", err)
	}
}

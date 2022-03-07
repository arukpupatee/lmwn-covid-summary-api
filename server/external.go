package server

import (
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type CovidCaseData struct {
	ConfirmDate    string
	Age            int
	Gender         string
	GenderEn       string // TODO: Refactor to enum
	Nation         string
	NationEn       string
	Province       string
	ProvinceEn     string
	ProvinceId     int
	District       string
	StatQuarantine int
}

type CovidCaseResponseBody struct {
	Data []CovidCaseData
}

func fetchCovidCase() (covidCaseResponseBody CovidCaseResponseBody, err error) {
	covidCaseResponseBody = CovidCaseResponseBody{}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	httpClient := &http.Client{Transport: tr}

	resp, err := httpClient.Get("https://static.wongnai.com/devinterview/covid-cases.json")

	if err != nil {
		return covidCaseResponseBody, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return covidCaseResponseBody, err
	}

	err = json.Unmarshal(body, &covidCaseResponseBody)

	if err != nil {
		return covidCaseResponseBody, err
	}

	return covidCaseResponseBody, nil
}

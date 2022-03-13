package models

type AgeGroup struct {
	Young   int `json:"0-30"`
	Middle  int `json:"31-60"`
	Old     int `json:"61+"`
	Unknown int `json:"N/A"`
}

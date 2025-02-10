package tools

import (
	"text/template"
)

var Tp *template.Template

type ErrorPage struct {
	Code         int
	ErrorMessage string
}

type Artists struct {
	Id           int          `json:"id"`
	Image        string       `json:"image"`
	Name         string       `json:"name"`
	Members      []string     `json:"members"`
	CreationDate int          `json:"creationDate"`
	FirstAlbum   string       `json:"firstAlbum"`
	Locations    string    `json:"locations"`
	ConcertDates string `json:"concertDates"`
	Relations    string    `json:"relations"`
}


type Locations struct {
	Id        int          `json:"id"`
	Locations []string     `json:"locations"`
	Dates     ConcertDates `json:"dates"`
}

type ConcertDates struct {
	Id    int      `json:"id"`
	Dates []string `json:"dates"`
}

type Relations struct {
	Id             int            `json:"id"`
	DatesLocations []ConcertDates `json:"datesLocations"`
}

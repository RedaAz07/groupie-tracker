package handler

import (
	"encoding/json"
	"net/http"

	"groupie/tools"
)

func Fetch_locations(url string) (tools.Locations, error) {
	var loc tools.Locations
	resp, err := http.Get(url)
	if err != nil {
		return loc, err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&loc)
	return loc, err
}
func Fetch_dates(url string) (tools.ConcertDates, error) {
	var dates tools.ConcertDates
	resp, err := http.Get(url)
	if err != nil {
		return dates, err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&dates)
	return dates, err
}
func Fetch_relation(url string) (tools.Relations, error) {
	var rel tools.Relations
	resp, err := http.Get(url)
	if err != nil {
		return rel, err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&rel)
	return rel, err
}
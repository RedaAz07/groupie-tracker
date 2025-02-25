package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"groupie/helpers"
	"groupie/tools"
)

func Detail_Card_Func(w http.ResponseWriter, r *http.Request) {
	// check the method
	if r.Method != http.MethodGet {
		// execute the status page template
		helpers.RenderTemplates(w, "statusPage.html", tools.ErrorMethodnotAll, http.StatusMethodNotAllowed)
		return
	}
	// fetching the artists
	var fetch_artist []tools.Artists
	url := "https://groupietrackers.herokuapp.com/api/artists"
	// get the api data
	res, err := http.Get(url)
	if err != nil {
		helpers.RenderTemplates(w, "statusPage.html", tools.ErrorInternalServerErr, http.StatusInternalServerError)
		return
	}
	defer res.Body.Close()
	// decode the jsone data
	err = json.NewDecoder(res.Body).Decode(&fetch_artist)
	if err != nil {
		helpers.RenderTemplates(w, "statusPage.html", tools.ErrorInternalServerErr, http.StatusInternalServerError)
		return
	}
	type fetchingData struct {
		Artist    *tools.Artists
		Locations *tools.Locations
		Dates     *tools.ConcertDates
		Relations *tools.Relations
	}
	// get the id from url
	id := r.URL.Query().Get("id")
	// to int
	Id, err := strconv.Atoi(id)
	if err != nil {
		helpers.RenderTemplates(w, "statusPage.html", tools.ErrorBadReq, http.StatusBadRequest)
		return
	}
	var artistFound *tools.Artists
	// get the user
	for i := range fetch_artist {
		if Id == fetch_artist[i].Id {
			artistFound = &fetch_artist[i]
			break
		}
	}
	//  to see if the user exists
	if artistFound == nil {
		helpers.RenderTemplates(w, "statusPage.html", tools.ErrorNotFound, http.StatusNotFound)
		return
	}
	var locations *tools.Locations
	var dates *tools.ConcertDates
	var relations *tools.Relations
	// declarations of urls
	Locations_url := artistFound.Locations
	ConcertDates_url := artistFound.ConcertDates
	Relations_url := artistFound.Relations
	
	// fetch the location and get the result  in the location variable
	errr := helpers.Fetch_By_Id(Locations_url, &locations)
	if errr != nil {
		helpers.RenderTemplates(w, "statusPage.html", tools.ErrorInternalServerErr, http.StatusInternalServerError)
		return
	}
	// fetch the dates and get the result  in the dates variable
	errr = helpers.Fetch_By_Id(ConcertDates_url, &dates)
	if errr != nil {
		helpers.RenderTemplates(w, "statusPage.html", tools.ErrorInternalServerErr, http.StatusInternalServerError)
		return
	}
	// fetch the relations and get the result  in the relations variable
	errr = helpers.Fetch_By_Id(Relations_url, &relations)
	if errr != nil {
		helpers.RenderTemplates(w, "statusPage.html", tools.ErrorInternalServerErr, http.StatusInternalServerError)
		return
	}
	// set all the data  that we found into the fetching variable
	fetching_data := fetchingData{
		Artist:    artistFound,
		Locations: locations,
		Dates:     dates,
		Relations: relations,
	}
	helpers.RenderTemplates(w, "detailsCard.html", fetching_data, http.StatusOK)
}

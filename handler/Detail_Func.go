package handler

import (
	"net/http"
	"strconv"

	"groupie/helpers"
	"groupie/tools"
)

func DetailFunc(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	Id, err := strconv.Atoi(id)
	if err != nil {
		helpers.RenderTemplates(w, "statusPage.html", nil, 400) 
		return
	}

	var searchArtist *tools.Artists
	for _, artist := range Artists {
		if Id == artist.Id {
			searchArtist = &artist
			break
		}
	}

	if searchArtist == nil {
		helpers.RenderTemplates(w, "statusPage.html", nil, 404) // Not Found
		return
	}

	helpers.RenderTemplates(w, "detailsCard.html", searchArtist, 200)
}

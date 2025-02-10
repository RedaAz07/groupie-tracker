package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"groupie/helpers"
	"groupie/tools"
)

func DetailFunc(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	Id, _ := strconv.Atoi(id)
	fmt.Println(Id)

	var searchArtist *tools.Artists
	for _, artist := range Artists {
		if Id == artist.Id {
			searchArtist = &artist
			break
		}

		if searchArtist == nil {
			helpers.RenderTemplates(w, "statusPage.html", nil, 404)
			return
		}
		helpers.RenderTemplates(w, "detailsCard.html", searchArtist, 200)
	}
}

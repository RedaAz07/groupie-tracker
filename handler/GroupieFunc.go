package handler

import (
	tools "groupie/Tools"
	"net/http"
)

func GroupieFunc(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		errore := tools.ErrorPage{
			Code:         http.StatusNotFound,
			ErrorMessage: "The page you are looking for might have been removed, had its name changed, or is temporarily unavailable.",
		}

		w.WriteHeader(http.StatusNotFound)
		// execute the not found  template
		tools.Tp.ExecuteTemplate(w, "statusPage.html", errore)
		return
	}
	tools.Tp.ExecuteTemplate(w, "index.html", nil)
}

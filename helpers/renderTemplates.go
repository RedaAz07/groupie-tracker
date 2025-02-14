package helpers

import (
	"bytes"
	"net/http"

	"groupie/tools"
)

func RenderTemplates(w http.ResponseWriter, temp string, data interface{}, status int) {
	var buf bytes.Buffer
	// exucut the template with buffer to chekc if there is an error in  our template
	err := tools.Tp.ExecuteTemplate(&buf, temp, data)
	if err != nil {
		errore := tools.ErrorInternalServerErr
		w.WriteHeader(http.StatusInternalServerError)
		tools.Tp.ExecuteTemplate(w, "statusPage.html", errore)
		return
	}
	w.WriteHeader(status)
	w.Write(buf.Bytes())
}

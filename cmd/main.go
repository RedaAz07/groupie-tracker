package main

import (
	"fmt"
	"net/http"
	"text/template"

	"groupie/handler"
	tools "groupie/tools"
)

func main() {
	var err error
	// parse all the html file from the template folder to variable Tp
	tools.Tp, err = template.ParseGlob("template/*.html")
	if err != nil {
		fmt.Println("Error parsing templates: ", err)
		return
	}

	// Register handlers
	http.HandleFunc("/styles/", handler.StyleFunc)
	http.HandleFunc("/", handler.GroupieFunc)
	fmt.Println("Server running at http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}

package rest

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os"
)

// SimplePage is
type SimplePage struct {
	Title  string
	Author string
}

func welcomeHandler() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode("successfully linked")
	}
}

func indexHandler() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		p := SimplePage{Title: "Why rexy is the best", Author: "Rexy"}
		tpl, err := template.ParseFiles("./pkg/http/rest/templates/index.html")
		if err != nil {
			fmt.Println("Couldn't parse html for index file! \n\n > Error:\n", err)
			os.Exit(http.StatusNotFound)
		} else {
			tpl.Execute(w, p)
		}
	}
}

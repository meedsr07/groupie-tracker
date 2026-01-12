package handlers

import (
	"html/template"
	"net/http"

	"groupie-tracker/models"
	"groupie-tracker/utils"
)

var artists []models.Artist

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		ErrorHandler(w,"method not allwed" ,405)
		return
	}
	// template.ParseFiles reades the html file when he found action like {{.}} he stocks in template object
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		ErrorHandler(w,"intarnal server error",500)
		return
	}

	artists, err = utils.FetchArtists()
	if err != nil {
		ErrorHandler(w,"intarnal server error",500)
		return
	}

	err = tmpl.Execute(w, artists)
	if err != nil {
		ErrorHandler(w,"intarnal server error",500)
		return
	}
}

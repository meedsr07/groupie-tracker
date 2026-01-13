package handlers

import (
	"bytes"
	"html/template"
	"net/http"

	"groupie-tracker/models"
	"groupie-tracker/utils"
)

var artists []models.Artist

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		ErrorHandler(w, "method not allowed", 405)
		return
	}
	if r.URL.Path != "/"  {
		ErrorHandler(w, "page not found", 404)
		return
	}
	// template.ParseFiles reades the html file when he found action like {{.}} he stocks in template object
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		ErrorHandler(w, "interanl server error", 500)
		return
	}

	artists, err = utils.FetchArtists()
	if err != nil {
		ErrorHandler(w, "interanl server error", 500)
		return
	}
	var buff bytes.Buffer
	err = tmpl.Execute(&buff, artists)
	if err != nil {
		ErrorHandler(w, "interanl server error", 500)
		return
	}
	w.Write(buff.Bytes())
}

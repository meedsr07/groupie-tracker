package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func main() {
	artists,err := FitchArtistData()
	if err !=nil{
		log.Fatal(err)
	}
	locations,err:=FitchLocationData()
	if err!=nil{
		log.Fatal(err)
	}
	http.HandleFunc("/", Handeler)
	http.HandleFunc("/artist", OneArtist)
	fmt.Println("Starting server on http://localhost:8080")
	Err := http.ListenAndServe(":8080", nil)
	if Err != nil {
		log.Fatal(Err)
	}
}

func Handeler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not alewed", 405)
		return
	}
	tmpl, err := template.ParseFiles("templetes/index.html")
	if err != nil {
		http.Error(w, "intrela erorr serevr", 500)
	}
	tmpl.Execute(w, artists)
}

func OneArtist(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not alewed", 405)
		return
	}
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Bad request", 400)
		return
	}
	artistId, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "bad request", 400)
		return
	}
	found := false
	var selectedArtist Artist
	for _, v := range artists {
		if v.ID == artistId {
			selectedArtist = v
			found = true
			break
		}
	}
	if !(found) {
		http.Error(w, "page not found", 404)
		return
	}
	tmpl, err := template.ParseFiles("templetes/artist.html")
	if err != nil {
		http.Error(w, "internal server error", 500)
		return
	}
	err = tmpl.Execute(w, selectedArtist)
	if err != nil {
		http.Error(w, "internal server error", 500)
		return
	}
}

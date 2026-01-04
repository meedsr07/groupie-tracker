package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"groupie/internal"
)

var (
	indexTmpl  = template.Must(template.ParseFiles("templates/index.html"))
	artistTmpl = template.Must(template.ParseFiles("templates/artist.html"))
)

// ArtistPage يربط الفنان بالـ Relation الخاص به
type ArtistPage struct {
	Artist   internal.Artist
	Relation internal.Relation
}

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/artist", artistHandler)

	log.Println("Starting server on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// indexHandler صفحة الرئيسية لعرض جميع الفنانين
func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	artists, err := internal.FetchArtists()
	if err != nil {
		http.Error(w, "Failed to fetch artists", http.StatusInternalServerError)
		return
	}

	if err := indexTmpl.Execute(w, artists); err != nil {
		http.Error(w, "Template error", http.StatusInternalServerError)
	}
}

// artistHandler صفحة تفاصيل الفنان + الحفلات
func artistHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.NotFound(w, r)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	artists, err := internal.FetchArtists()
	if err != nil {
		http.Error(w, "Failed to fetch artists", http.StatusInternalServerError)
		return
	}

	relations, err := internal.FetchRelations()
	if err != nil {
		http.Error(w, "Failed to fetch relations", http.StatusInternalServerError)
		return
	}

	var page ArtistPage

	for _, artist := range artists {
		if artist.ID == id {
			page.Artist = artist
			break
		}
	}

	for _, rel := range relations {
		if rel.ID == id {
			page.Relation = rel
			break
		}
	}

	if page.Artist.ID == 0 {
		http.NotFound(w, r)
		return
	}

	if err := artistTmpl.Execute(w, page); err != nil {
		http.Error(w, "Template error", http.StatusInternalServerError)
	}
}

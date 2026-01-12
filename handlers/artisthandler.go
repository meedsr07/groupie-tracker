package handlers

import (
	"html/template"
	"net/http"
	"strconv"

	"groupie-tracker/models"
	"groupie-tracker/utils"
)

func ArtistHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		ErrorHandler(w,"method not allewed",405)
		return
	}
	id := r.URL.Query().Get("id")
	if id == "" {
		ErrorHandler(w,"page not fuond",404)
		return
	}
	artistId, err := strconv.Atoi(id)
	if err != nil {
		ErrorHandler(w,"page not fuond",404)
		return
	}
	found := false
	var selectedArtist models.Artist
	for _, v := range artists {
		if v.ID == artistId {
			selectedArtist = v
			found = true
			break
		}
	}
	if !(found) {
		ErrorHandler(w,"page not found",404)
		return
	}

	location, err := utils.FetchLocations()
	if err != nil {
		ErrorHandler(w,"internal server error",500)
		return
	}
	artistLocation:= GetArtistLocation(location, artistId)

	date, err := utils.FetchDates()
	if err != nil {
		ErrorHandler(w,"internal server error",500)
		return
	}
	artistdate:= GetArtistDate(date, artistId)
	artistrelation,err:=utils.FetchRelations()
	if err != nil{
		ErrorHandler(w,"intarnal server error",500)
		return
	}
	datarelation:=GetArtidtRelation(artistrelation,artistId)

	data := models.ArtistPageData{
		Artist:   selectedArtist,
		Location: artistLocation,
		Date:     artistdate,
		DatesLocations: datarelation.DatesLocations,
	}

	tmpl, err := template.ParseFiles("templates/artist.html")
	if err != nil {
		ErrorHandler(w,"intarnal server error",500)
		return
	}
	err =tmpl.Execute(w, data)
	if err != nil {
		ErrorHandler(w,"intarnal server error",500)
		return
	}
}

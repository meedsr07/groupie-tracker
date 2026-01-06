package main

import (
	"fmt"
	"log"
	"net/http"

	
)

type Artist struct {
	ID           int      `json:"id"`
	Name         string   `json:"name"`
	Image        string   `json:"image"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
}

var artists []Artist

func main() {
	url := "https://groupietrackers.herokuapp.com/api/artists"
	data, err := GetAPIData(url)
	if err != nil {
		log.Fatal(err)
	}
	artists, err = DecodeData(data)
	if err != nil {
		log.Fatal(err)
	}
	http.HandleFunc("/", Handeler)
	fmt.Println("Starting server on http://localhost:8080")
	Err := http.ListenAndServe(":8080", nil)
	if Err != nil {
		log.Fatal(err)
	}
}

func Handeler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not alewed", 405)
		return
	}
	fmt.Fprintf(w, "Artists count: %d\n", len(artists))
	if len(artists) > 0 {
		fmt.Fprintf(w, "First artist: %s\n", artists[0].Name)
	}
}

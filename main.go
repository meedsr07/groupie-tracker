package main

import (
	"fmt"
	"log"
	"net/http"
	"html/template"
)

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
	tmpl,err := template.ParseFiles("templetes/index.html")
	if err != nil{
		http.Error(w,"intela erorr serevr",500)
	}
	tmpl.Execute(w,artists)

}

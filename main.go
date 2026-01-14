package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"groupie-tracker/handlers"
)

func main() {
	// قراءة PORT من environment
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	http.ListenAndServe(":"+port, nil)

	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/artist", handlers.ArtistHandler)
	http.HandleFunc("/static/", handlers.StaticHandler)

	fmt.Printf("Server successfully started at http://localhost:%s\n", port)

	// Listen على PORT الصحيح
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}

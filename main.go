package main

import (
    "fmt"
    "log"
    "net/http"
    "os"

    "groupie-tracker/handlers"
)

func main() {
    port := os.Getenv("PORT") // قراءة البورت من environment
    if port == "" {
        port = "8080" // لو محلياً
    }

    http.HandleFunc("/", handlers.HomeHandler)
    http.HandleFunc("/artist", handlers.ArtistHandler)
    http.HandleFunc("/static/", handlers.StaticHandler)

    fmt.Printf("Server successfully started at http://localhost:%s\n", port)

    log.Fatal(http.ListenAndServe(":"+port, nil)) // Listen على PORT الصحيح
}

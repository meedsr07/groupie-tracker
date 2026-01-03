package main

import (
	"fmt"
	"log"
	"net/http"
)
func Handler(w http.ResponseWriter, r *http.Request) {
 if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	fs := http.FileServer(http.Dir("./static"))
	fs.ServeHTTP(w, r)	
}
func main() {
	http.HandleFunc("/", Handler)
	
	fmt.Println("Starting server in : http://localhost:8080")
	err  := http.ListenAndServe(":8080",nil)
	if err != nil {
		log.Fatal(err)
	}

}
package handlers

import (
	"net/http"
	"os"


)

func StaticHandler(w http.ResponseWriter, r *http.Request) {
	filePath := r.URL.Path[1:]

	fileInfo, err := os.Stat(filePath)
	if os.IsNotExist(err) || fileInfo.IsDir() {
		ErrorHandler(w, "Resource Not Found", http.StatusNotFound)
		return
	}

	staticServer := http.StripPrefix("/static/", http.FileServer(http.Dir("static")))
	staticServer.ServeHTTP(w, r)
}

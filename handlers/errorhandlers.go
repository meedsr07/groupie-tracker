package handlers

import (
	"html/template"
	"net/http"
)

type ErrorPage struct {
	Code int
	Message string
}
func ErrorHandler(w http.ResponseWriter, ErrorMessag string, statusCode int) {

	errorPage := ErrorPage{
		Code: statusCode,
		Message: ErrorMessag,
	}

	tmpl, err := template.ParseFiles("templates/error.html")
	if err != nil {
		ErrorHandler(w, "internal server error", 500)
		return
	}

	tmpl.Execute(w, errorPage)
}

package handlers

import (
	"net/http"
)

func Index(writer http.ResponseWriter, request *http.Request) {
	generateHTML(writer, nil, "layout", "auth.navbar", "index")
}


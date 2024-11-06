package routes

import (
	"encoding/json"
	"net/http"
)

func HandleMessage(w http.ResponseWriter, r *http.Request) {
	payload := "{ \"message\": \"Hello World!\" }"

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(payload)
}

func HandleWebPages(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	filePath := "./public/static" + path

	if path == "/" {
		filePath += "index"
	}

	filePath += ".html"

	http.ServeFile(w, r, filePath)
}

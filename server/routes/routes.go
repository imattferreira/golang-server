package routes

import (
	"encoding/json"
	"net/http"
)

func GetMessage(w http.ResponseWriter, r *http.Request) {
	payload := "{ \"message\": \"Hello World!\" }"

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(payload)
}

func GetWebPages() http.Handler {
	return http.FileServer(http.Dir("./public/static"))
}

package routes

import (
	"io"
	"net/http"
)

func GetMessage(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "{ \"message\": \"Hello World!\" }")
}

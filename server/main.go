package main

import (
	"distributed-system/server/routes"
	"fmt"
	"net/http"
)

func registerRoutes() {
	http.HandleFunc("/", routes.HandleWebPages)
	http.HandleFunc("/api/message", routes.HandleMessage)
}

func startServer() {
	err := http.ListenAndServe(":3000", nil)

	if err != nil {
		panic("error starting server")
	}
}

func main() {
	registerRoutes()
	fmt.Println("🚀 Server running at http://localhost:3000")
	startServer()
}

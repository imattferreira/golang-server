package main

import (
	"distributed-system/server/routes"
	"fmt"
	"net/http"
)

func registerRoutes() {
	http.Handle("/", routes.GetWebPages())
	http.HandleFunc("/message", routes.GetMessage)
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

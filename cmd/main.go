package main

import (
	"log"
	"net/http"
	"tasks/internal/handlers"
)

func main() {
	server := handlers.NewTaskServer()

	mux := http.NewServeMux()
	mux.HandleFunc("/", server.HandlerTask)
	log.Fatal(http.ListenAndServe("localhost:2022", mux))
}

package main

import (
	"io"
	"net/http"
	"time"
	"log"
)

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/hello", hello)

	s := &http.Server{
		Addr:		":8080",
		Handler:	mux,
		ReadTimeout:	2 * time.Second,
		WriteTimeout: 	2 * time.Second,
		IdleTimeout:	5 * time.Second,
	}

	log.Println("Starting server :8080")
	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal("Server startup failed")
	}
}

func hello (w http.ResponseWriter, req *http.Request) {

	io.WriteString(w, "Hello, World!!!!")
}



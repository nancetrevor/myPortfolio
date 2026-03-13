package main

import (
	"log"
	"net/http"
	"time"

	"github.com/lfizzikz/portfolioSite/internal/handlers"
)

func main() {
	mux := http.NewServeMux()

	fs := http.FileServer(http.Dir("web/static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})

	mux.HandleFunc("/", handlers.HomeHandler)
	mux.HandleFunc("/terminal/", handlers.TerminalHandler)

	serverHost := "127.0.0.1:8080"
	httpServer := &http.Server{
		Addr:         serverHost,
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Printf("Server listening on %s", httpServer.Addr)
	log.Fatal(httpServer.ListenAndServe())
}

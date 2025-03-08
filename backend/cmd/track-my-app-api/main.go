package main

import (
	"log"
	"net/http"
	"track-my-app-api/db"

	"github.com/rs/cors"
)

func main() {
	defer func() {
		if err := db.Client.Disconnect(db.Ctx); err != nil {
			log.Fatal("Error while disconnecting client: ", err)
		}
		db.Cancel() // Release the context
		log.Println("Shutdown cleanup complete")
	}()

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world"))
	})
	handler := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:5173"},
		AllowedMethods: []string{"GET", "POST"},
	}).Handler(mux)

	log.Println("Starting server on :8080")
	err := http.ListenAndServe(":8080", handler)
	if err != nil {
		log.Fatal(err)
	}
}

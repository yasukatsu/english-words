package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	pos "y_ara/english-words/backend/internal/app/postgres"
)

func main() {
	allowedOrigins := handlers.AllowedOrigins([]string{"http://localhost:8080"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT"})

	r := mux.NewRouter()
	r.HandleFunc("/list", listWords)

	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(allowedOrigins, allowedMethods)(r)))
}

func listWords(w http.ResponseWriter, r *http.Request) {
	words, err := pos.ListWords()

	if err != nil {
		log.Fatalf("ListWords err is %v\n", err)
		return
	}

	byteWords, err := json.Marshal(words)
	if err != nil {
		log.Fatalf("json.Marshal err is %v\n", err)
		return
	}
	w.Write(byteWords)
}

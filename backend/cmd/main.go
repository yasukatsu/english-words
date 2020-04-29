package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	pos "y_ara/english-words/backend/internal/app/postgres"
	"y_ara/english-words/backend/internal/pkg/config"
)

func main() {
	e := config.GetGoEnv()
	err := config.NewConfig(e)
	if err != nil {
		log.Fatalf("load config %v", err)
		return
	}
	allowedOrigins := handlers.AllowedOrigins([]string{"http://localhost:8080"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT"})

	r := mux.NewRouter()
	r.HandleFunc("/list", listWords)

	// words, err := pos.ListWords()
	// if err != nil {
	// 	log.Println("ListWords err", err)
	// 	return
	// }
	// fmt.Println(words)

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

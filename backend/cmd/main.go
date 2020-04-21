package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	pos "y_ara/english-words/backend/internal/app/postgres"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// Words ...
type Words []struct {
	Meanings          string `json:"Meanings"`
	EnglishDefinition string `json:"EnglishDefinition"`
	JTranslation      string `json:"JTranslation"`
}

var (
	path  = "./resourse/datas.json"
	pathA = "./resourse/a.json"
	pathB = "./resourse/b.json"
	pathC = "./resourse/c.json"
	pathD = "./resourse/d.json"
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
	fmt.Println(words)

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

// func list(w http.ResponseWriter, r *http.Request) {
// 	body, err := readFile(path)
// 	if err != nil {
// 		log.Fatalf("ReadFile err is %v", err)
// 		return
// 	}

// 	w.Write(body)
// }

// func readFile(path string) ([]byte, error) {
// 	raw, err := ioutil.ReadFile(path)
// 	if err != nil {
// 		log.Fatalf("ReadFile err is %v", err)
// 		return nil, err
// 	}
// 	body := bytes.TrimPrefix(raw, []byte("\xef\xbb\xbf"))
// 	return body, nil
// }

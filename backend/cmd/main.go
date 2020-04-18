package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"

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
	r.HandleFunc("/list", list)

	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(allowedOrigins, allowedMethods)(r)))
}

func list(w http.ResponseWriter, r *http.Request) {
	body, err := readFile(path)
	if err != nil {
		log.Fatalf("ReadFile err is %v", err)
		return
	}

	w.Write(body)
}

func readFile(path string) ([]byte, error) {
	raw, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("ReadFile err is %v", err)
		return nil, err
	}
	body := bytes.TrimPrefix(raw, []byte("\xef\xbb\xbf"))
	return body, nil
}

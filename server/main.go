package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// English ...
type English struct {
	Words []struct {
		Meanings          string `json:"Meanings"`
		EnglishDefinition string `json:"EnglishDefinition"`
		POS               string `json:"POS"`
		JTranslation      string `json:"J Translation"`
	} `json:"words"`
}

var (
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
	body, err := readFile(pathA)
	if err != nil {
		log.Fatalf("ReadFile err is %v", err)
		return
	}

	var e English
	err = json.Unmarshal(body, &e)
	if err != nil {
		log.Fatalf("Unmarshal err is %v", err)
		return
	}

	// fmt.Println(&e)
	response := fmt.Sprint(&e)

	w.Write([]byte(response))
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

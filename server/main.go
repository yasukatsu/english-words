package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

// English ...
type English struct {
	Words []struct {
		Meanings          string `json:"Meanings"`
		EnglishDefinition string `json:"English Definition"`
		POS               string `json:"POS"`
		JTranslation      string `json:"J Translation"`
	} `json:"words"`
}

var path = "./resourse/data.json"

func main() {
	body, err := readFile(path)
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

	fmt.Println(&e)
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

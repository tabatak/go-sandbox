package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

type Payload struct {
	Content string
}

func main() {

	payload := Payload{
		Content: "1\n2\n3\r\n",
	}

	encoded, err := json.Marshal(payload)
	if err != nil {
		log.Fatal(err)
	}

	req, err := http.NewRequest(
		"POST",
		"http://127.0.0.1:8080/",
		bytes.NewBuffer(encoded),
	)
	if err != nil {
		log.Fatal(err)
	}

	// Content-Type 設定
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	client.Do(req)
}

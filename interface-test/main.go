package main

import (
	"log"

	"github.com/tabatak/go-sandbox/interface-test/client"
)

func main() {
	client := &client.APIClientImpl{}

	got := client.Get("main key")
	log.Printf("got=%s", got)
}

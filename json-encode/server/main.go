package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		b, _ := ioutil.ReadAll(r.Body)
		log.Printf("body:\n%s", string(b))
		var decoded map[string]string
		json.Unmarshal(b, &decoded)
		log.Printf("Content:%s", decoded["Content"])
	})
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))

}

package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
)

func handler(w http.ResponseWriter, r *http.Request) {
	dump, err := httputil.DumpRequest(r, true)
	if err != nil {
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
		return
	}
	fmt.Println(string(dump))

	res := `{
		"data": {
		  "type": "gif",
		  "id": "l2JhxfHWMBWuDMIpi",
		  "title": "cat love GIF by The Secret Life Of Pets",
		  "image_url": "https://media1.giphy.com/media/l2JhxfHWMBWuDMIpi/giphy.gif",
		  "caption": "",
		},
		"meta": {
		  "status": 200,
		  "msg": "OK",
		  "response_id": "5b105e44316d3571456c18b3"
		}
	  }`
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(res))
	// fmt.Fprintf(w, res)
}

func main() {
	var httpServer http.Server
	http.HandleFunc("/", handler)
	log.Println("start http listening :18888")
	httpServer.Addr = ":18888"
	log.Println(httpServer.ListenAndServe())
}

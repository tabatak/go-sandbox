package main

import (
	"log"
	"os"

	"github.com/tabatak/go-sandbox/httptest-totorial/stationdata"
)

func main() {
	// ドメインを指定
	os.Setenv("STATIONDATA_API_DOMAIN", "http://www.ekidata.jp")

	res, err := stationdata.GetLineInfoFromPrefCode(100)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("response=%s", res)
}

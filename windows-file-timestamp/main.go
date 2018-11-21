package main

import (
	"log"
	"os"
	"time"
)

func main() {
	file, err := os.Open("main.go")
	if err != nil {
		log.Fatal(err)
	}

	fileInfo, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("ModTime=%s (%d)", fileInfo.ModTime().Format(time.RFC3339), fileInfo.ModTime().Unix())

	name, offset := time.Now().Zone()
	log.Printf("timezone: name=%s, offset=%d", name, offset)
}

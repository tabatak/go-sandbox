package main

import (
	"log"
	"path/filepath"
)

func main() {
	// *.cfg
	fileNames, err := filepath.Glob("*.cfg")
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("pattern: %s\n", "*.cfg")
	for _, fileName := range fileNames {
		log.Printf("\tfile name: %s", fileName)
	}

	// *test*
	fileNames, err = filepath.Glob("*test*")
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("pattern: %s\n", "*test*")
	for _, fileName := range fileNames {
		log.Printf("\tfile name: %s", fileName)
	}
}

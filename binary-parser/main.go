package main

import (
	"log"

	"github.com/tabatak/go-sandbox/binary-parser/parser"
)

func main() {

	data := []byte{0x41, 0x42, 0x43, 0x44, 0x30, 0x31, 0xFF, 0xFF}
	format := "val1:char:4 val2:char:2 val3:uint:2"

	parsedData := parser.Parse(data, format)
	log.Printf("%v", parsedData)
}

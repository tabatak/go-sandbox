package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Printf("go version: %s\n", runtime.Version())
}

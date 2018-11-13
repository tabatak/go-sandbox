package main

import (
	"time"
)

var a string
var done bool

func setup() {
	a = "hello, world"
	time.Sleep(100 * time.Second)
	done = true
}

func main() {
	go setup()
	i := 0
	for !done {
		// println(i)
		i++
	}
	println(a)
	println(i)
}

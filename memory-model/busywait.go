package main

var a string
var done bool

func setup() {
	a = "hello, world"
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

package main

var a, b int

func f() {
	a = 1
	b = 2
}

func g() {
	println(b)
	println(a)
}

func main() {
	go f()
	g()
}

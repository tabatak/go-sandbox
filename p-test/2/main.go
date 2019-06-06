package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	s := []string{"u", "d", "c", "s", "t", "b"}

	makeStr(n, s, "")
}

func makeStr(n int, s []string, tmp string) {
	fmt.Println(tmp)
	if n == len(tmp) {
		// fmt.Println(tmp)
	} else {
		for i := 0; i < len(s); i++ {
			makeStr(n, s, tmp+string(s[i]))
		}
	}
}

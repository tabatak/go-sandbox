package main

import (
	"fmt"
	"strings"
)

func main() {
	var n int
	fmt.Scan(&n)

	s := []string{"u", "d", "c", "s", "t", "b"}

	makeStr(n, s, "", "uud")
}

func makeStr(n int, s []string, tmp string, subsr string) {
	if n == len(tmp) {
		if strings.Contains(tmp, subsr) {
			fmt.Println(tmp)
		}
	} else {
		for i := 0; i < len(s); i++ {
			makeStr(n, s, tmp+string(s[i]), subsr)
		}
	}
}

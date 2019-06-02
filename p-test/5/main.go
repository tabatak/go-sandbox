package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	ss := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&ss[i])
	}

	for i := 0; i < n; i++ {
		s := ss[i]
		changed := false

		if len(s) > 0 {
			if s[len(s)-1:len(s)] == "s" ||
				(len(s) > 1 && s[len(s)-2:len(s)] == "sh") ||
				(len(s) > 1 && s[len(s)-2:len(s)] == "ch") ||
				s[len(s)-1:len(s)] == "o" ||
				s[len(s)-1:len(s)] == "x" {
				changed = true
				fmt.Println(fmt.Sprintf("%ses", s))
			} else if s[len(s)-1:len(s)] == "f" {
				changed = true
				fmt.Println(fmt.Sprintf("%sves", s[0:len(s)-1]))
			} else if len(s) > 1 && s[len(s)-2:len(s)] == "fe" {
				changed = true
				fmt.Println(fmt.Sprintf("%sves", s[0:len(s)-2]))
			} else if s[len(s)-1:len(s)] == "y" {
				if len(s) < 2 {
					changed = true
					fmt.Println(fmt.Sprintf("%sies", s[0:len(s)-1]))
				} else if s[len(s)-2:len(s)-1] != "a" &&
					s[len(s)-2:len(s)-1] != "i" &&
					s[len(s)-2:len(s)-1] != "u" &&
					s[len(s)-2:len(s)-1] != "e" &&
					s[len(s)-2:len(s)-1] != "o" {
					changed = true
					fmt.Println(fmt.Sprintf("%sies", s[0:len(s)-1]))
				}
			}
		}

		if !changed {
			fmt.Println(fmt.Sprintf("%ss", s))
		}

	}
}

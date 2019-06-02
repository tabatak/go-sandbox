package main

import "fmt"

func main() {
	s := []string{"u", "d", "c", "s", "t", "b"}
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s); j++ {
			for k := 0; k < len(s); k++ {
				fmt.Println(s[i] + s[j] + s[k])
			}
		}
	}
}

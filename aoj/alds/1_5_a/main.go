package main

import (
	"fmt"
)

var (
	n int
	a []int
)

func main() {
	fmt.Scan(&n)
	a = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	var q int
	fmt.Scan(&q)
	ms := make([]int, q)
	for i := 0; i < q; i++ {
		fmt.Scan(&ms[i])
	}

	//回答
	for i := 0; i < q; i++ {
		res := solve(0, ms[i])
		if res {
			fmt.Println("yes")
		} else {
			fmt.Println("no")
		}
	}

}

func solve(i, m int) bool {
	if m == 0 {
		return true
	}

	if i >= n {
		return false
	}

	res := solve(i+1, m) || solve(i+1, m-a[i])
	return res
}

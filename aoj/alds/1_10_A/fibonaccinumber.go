package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	f := make([]int, n+1)

	fmt.Println(fib(n, f))

	fibSlice := makeFibonacci(n)
	fmt.Println(fibSlice)
	fmt.Println(fibSlice[n])
}

// メモ化再帰
func fib(n int, f []int) int {
	if n == 0 {
		f[n] = 1
		return f[n]
	}
	if n == 1 {
		f[n] = 1
		return 1
	}

	if f[n] != 0 {
		return f[n]
	}

	f[n] = fib(n-2, f) + fib(n-1, f)
	return f[n]
}

// 動的計画法
func makeFibonacci(n int) []int {
	f := make([]int, n+1)

	f[0] = 1
	f[1] = 1

	for i := 2; i <= n; i++ {
		f[i] = f[i-2] + f[i-1]
	}

	return f
}

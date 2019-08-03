package main

import "fmt"

type Node struct {
	p int
	l int
	r int
}

const (
	NIL = -1
	MAX = 100005
)

var T [MAX]Node
var n int
var D [MAX]int

func main() {
	var d, v, c, l, r int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		T[i].p = NIL
		T[i].l = NIL
		T[i].r = NIL
	}

	for i := 0; i < n; i++ {
		fmt.Scan(&v)
		fmt.Scan(&d)

		for j := 0; j < d; j++ {
			fmt.Scan(&c)
			if j == 0 {
				T[v].l = c
			} else {
				T[l].r = c
			}
			l = c
			T[c].p = v
		}
	}

	for i := 0; i < n; i++ {
		if T[i].p == NIL {
			r = i
		}
	}

	rec(r, 0)

	for i := 0; i < n; i++ {
		print(i)
	}

	return
}

func rec(u, p int) {
	D[u] = p
	if T[u].r != NIL {
		rec(T[u].r, p)
	}
	if T[u].l != NIL {
		rec(T[u].l, p+1)
	}
}

func print(u int) {
	format := "node %d: parent = %d, depth = %d, %s, %s\n"

	nodeType := ""
	if T[u].p == NIL {
		nodeType = "root"
	} else if T[u].l == NIL {
		nodeType = "leaf"
	} else {
		nodeType = "internal node"
	}

	children := "["
	c := T[u].l
	count := 0
	for c != NIL {
		if count > 0 {
			children += ", "
		}
		children += fmt.Sprintf("%d", c)

		c = T[c].r
		count++
	}
	children += "]"

	fmt.Printf(format, u, T[u].p, D[u], nodeType, children)
}

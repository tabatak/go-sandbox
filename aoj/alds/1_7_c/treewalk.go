package main

import (
	"fmt"
)

type Node struct {
	p, l, r int
}

const NIL = -1

var n, id, l, r int

func main() {
	fmt.Scan(&n)

	nodes := make([]Node, n)
	for i := 0; i < n; i++ {
		nodes[i].p = NIL
	}

	for i := 0; i < n; i++ {
		fmt.Scan(&id)
		fmt.Scan(&l)
		fmt.Scan(&r)

		nodes[id].l = l
		nodes[id].r = r

		if l != NIL {
			nodes[l].p = id
		}
		if r != NIL {
			nodes[r].p = id
		}
	}

	//nodesができたのでrootを探す
	root := -1
	for i := 0; i < n; i++ {
		if nodes[i].p == NIL {
			root = i
			break
		}
	}

	//Preorder
	fmt.Println("Preorder")
	preParse(nodes, root)
	fmt.Println("")
	//Inorder
	fmt.Println("Inorder")
	inParse(nodes, root)
	fmt.Println("")
	//Postorder
	fmt.Println("Postorder")
	postParse(nodes, root)
	fmt.Println("")

}

func preParse(nodes []Node, u int) {
	if u == NIL {
		return
	}

	fmt.Printf(" %d", u)
	preParse(nodes, nodes[u].l)
	preParse(nodes, nodes[u].r)
}

func inParse(nodes []Node, u int) {
	if u == NIL {
		return
	}

	inParse(nodes, nodes[u].l)
	fmt.Printf(" %d", u)
	inParse(nodes, nodes[u].r)
}

func postParse(nodes []Node, u int) {
	if u == NIL {
		return
	}

	postParse(nodes, nodes[u].l)
	postParse(nodes, nodes[u].r)
	fmt.Printf(" %d", u)
}

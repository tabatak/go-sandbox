package main

import (
	"fmt"
)

type Node struct {
	p, l, r int
	d       int
	h       int
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

	//depthを下に向かって設定
	setDepth(nodes, root, 0)

	//高さを設定
	for i := 0; i < n; i++ {
		setHeight(nodes, i)
	}

	//出力
	for i := 0; i < n; i++ {
		print(nodes, i)
	}
}

func setDepth(nodes []Node, current int, depth int) {
	if current == NIL {
		return
	}

	nodes[current].d = depth

	setDepth(nodes, nodes[current].l, depth+1)
	setDepth(nodes, nodes[current].r, depth+1)
}

func setHeight(nodes []Node, u int) int {
	h1 := 0
	h2 := 0

	if nodes[u].l != NIL {
		h1 = setHeight(nodes, nodes[u].l) + 1
	}
	if nodes[u].r != NIL {
		h2 = setHeight(nodes, nodes[u].r) + 1
	}

	h := h1
	if h1 < h2 {
		h = h2
	}
	nodes[u].h = h
	return h
}

func getSibling(nodes []Node, u int) int {
	if nodes[u].p == NIL {
		return NIL
	}

	if nodes[nodes[u].p].l != u && nodes[nodes[u].p].l != NIL {
		return nodes[nodes[u].p].l
	}
	if nodes[nodes[u].p].r != u && nodes[nodes[u].p].r != NIL {
		return nodes[nodes[u].p].r
	}
	return NIL
}

func print(nodes []Node, u int) {
	format := "node %d: parent = %d, sibling = %d, degree = %d, depth = %d, height = %d, %s\n"

	nodeType := ""
	if nodes[u].p == NIL {
		nodeType = "root"
	} else if nodes[u].l == NIL && nodes[u].r == NIL {
		nodeType = "leaf"
	} else {
		nodeType = "internal node"
	}

	sibling := getSibling(nodes, u)
	degree := 0
	if nodes[u].l != NIL {
		degree++
	}
	if nodes[u].r != NIL {
		degree++
	}

	fmt.Printf(format, u, nodes[u].p, sibling, degree, nodes[u].d, nodes[u].h, nodeType)

}

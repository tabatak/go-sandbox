package main

import (
	"fmt"
	"strings"
)

type Node struct {
	parent   int
	children []int
	depth    int
}

const NIL = -1

var n, id, k, c int

func main() {
	fmt.Scan(&n)

	nodes := make([]Node, n)

	for i := 0; i < n; i++ {
		nodes[i].parent = NIL
	}

	for i := 0; i < n; i++ {
		fmt.Scan(&id)
		fmt.Scan(&k)

		for j := 0; j < k; j++ {
			fmt.Scan(&c)
			nodes[id].children = append(nodes[id].children, c)
			nodes[c].parent = id
		}
	}

	//nodesができたのでrootを探す
	root := -1
	for i := 0; i < n; i++ {
		if nodes[i].parent == NIL {
			root = i
			break
		}
	}

	//depthを設定する
	rec(nodes, root, 0)

	//出力
	for i := 0; i < n; i++ {
		print(nodes, i)
	}

}

func rec(nodes []Node, currentNode int, depth int) {
	nodes[currentNode].depth = depth

	//現在のnodeのdepthに+1したものを子に配る（を繰り返す）
	for i := 0; i < len(nodes[currentNode].children); i++ {
		rec(nodes, nodes[currentNode].children[i], depth+1)
	}
}

func print(nodes []Node, u int) {
	format := "node %d: parent = %d, depth = %d, %s, %s\n"

	nodeType := ""
	if nodes[u].parent == NIL {
		nodeType = "root"
	} else if len(nodes[u].children) == 0 {
		nodeType = "leaf"
	} else {
		nodeType = "internal node"
	}

	// children := "["
	// for i := 0; i < len(nodes[u].children); i++ {
	// 	if len(children) > 1 {
	// 		children += ", "
	// 	}
	// 	children += fmt.Sprintf("%d", nodes[u].children[i])
	// }
	// children += "]"
	children := strings.Replace(fmt.Sprintf("%v", nodes[u].children), " ", ", ", -1)

	fmt.Printf(format, u, nodes[u].parent, nodes[u].depth, nodeType, children)
}

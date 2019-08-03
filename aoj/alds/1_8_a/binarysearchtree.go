package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Node struct {
	key                 int
	parent, left, right *Node
}

type BinarySearchTree struct {
	root *Node
	n    int
}

func (bst *BinarySearchTree) insert(x int) bool {
	p := bst.findLast(x)
	u := &Node{key: x}
	return bst.addChild(p, u)
}

func (bst *BinarySearchTree) findLast(x int) *Node {
	w := bst.root
	var prev *Node

	for w != nil {
		prev = w
		if x < w.key {
			w = w.left
		} else {
			w = w.right
		}
	}
	return prev
}

func (bst *BinarySearchTree) addChild(p, u *Node) bool {
	if p == nil {
		bst.root = u
	} else {
		if u.key < p.key {
			p.left = u
		} else {
			p.right = u
		}
		u.parent = p
	}
	bst.n++
	return true
}

func preorder(u *Node, keys []int) []int {
	if u != nil {
		keys = append(keys, u.key)
		keys = preorder(u.left, keys)
		keys = preorder(u.right, keys)
	}
	return keys
}

func inorder(u *Node, keys []int) []int {
	if u != nil {
		keys = inorder(u.left, keys)
		keys = append(keys, u.key)
		keys = inorder(u.right, keys)
	}
	return keys
}

func print(keys []int) {
	w := bufio.NewWriter(os.Stdout)
	for _, key := range keys {
		w.Write([]byte(fmt.Sprintf(" %d", key)))
	}
	w.Write([]byte("\n"))
	w.Flush()
}

func main() {
	var n int
	fmt.Scan(&n)

	bst := new(BinarySearchTree)

	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	for i := 0; i < n; i++ {
		sc.Scan()
		command := sc.Text()

		if command == "insert" {
			sc.Scan()
			x, _ := strconv.Atoi(sc.Text())
			//insert
			bst.insert(x)
		} else if command == "print" {
			//print
			print(inorder(bst.root, []int{}))
			print(preorder(bst.root, []int{}))
		}
	}

	return
}

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

func (bst *BinarySearchTree) find(x int) *Node {
	u := bst.root

	for u != nil && x != u.key {
		if x < u.key {
			u = u.left
		} else {
			u = u.right
		}
	}
	return u
}

func (bst *BinarySearchTree) Splice(u *Node) {
	var s, p *Node

	if u.left != nil {
		s = u.left
	} else {
		s = u.right
	}

	if u == bst.root {
		bst.root = s
		p = nil
	} else {
		p = u.parent
		if p.left == u {
			p.left = s
		} else {
			p.right = s
		}
	}
	if s != nil {
		s.parent = p
	}
	bst.n--
}

func (bst *BinarySearchTree) Remove(u *Node) {
	if u.left == nil || u.right == nil {
		//削除対象に片方の枝しかないまたは枝がない場合は、parentとその枝のノードをつなげば完了
		bst.Splice(u)
		u = nil
	} else {
		//両方枝がある場合は、自分より大きい値の中の最小値のノードで置き換える
		w := u.right
		for w.left != nil {
			w = w.left
		}
		u.key = w.key //削除対象だったノードの値を置き換える
		bst.Splice(w) //置き換えたノードを削除する
		w = nil
	}
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
		} else if command == "find" {
			sc.Scan()
			x, _ := strconv.Atoi(sc.Text())
			//find
			u := bst.find(x)
			if u != nil {
				fmt.Println("yes")
			} else {
				fmt.Println("no")
			}
		} else if command == "delete" {
			sc.Scan()
			x, _ := strconv.Atoi(sc.Text())
			//delete
			u := bst.find(x)
			if u != nil {
				bst.Remove(u)
			}

		}
	}

	return
}

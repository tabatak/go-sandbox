package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	w := bufio.NewWriter(os.Stdout)
	bh := NewBinaryHeap()

	for {
		sc.Scan()
		command := sc.Text()

		if command == "insert" {
			sc.Scan()
			x, _ := strconv.Atoi(sc.Text())
			//insert
			bh.Add(x)

		} else if command == "extract" {
			//remove
			x := bh.Remove()
			w.Write([]byte(fmt.Sprintf("%d\n", x)))

		} else if command == "end" {
			break
		}
	}

	w.Flush()
}

type BinaryHeap struct {
	n int
	a []int
}

func NewBinaryHeap() *BinaryHeap {
	return &BinaryHeap{}
}
func (bh *BinaryHeap) left(i int) int {
	return 2*i + 1
}

func (bh *BinaryHeap) right(i int) int {
	return 2*i + 2
}

func (bh *BinaryHeap) parent(i int) int {
	return (i - 1) / 2
}

func (bh *BinaryHeap) Add(x int) bool {
	if bh.n+1 > len(bh.a) {
		bh.resize()
	}

	bh.a[bh.n] = x
	bh.n++
	bh.bubbleUp(bh.n - 1)

	return true
}

func (bh *BinaryHeap) resize() {
	capNew := 1
	if 2*bh.n > 1 {
		capNew = 2 * bh.n
	}
	newArray := make([]int, capNew)
	copy(newArray, bh.a)
	bh.a = newArray
}

func (bh *BinaryHeap) bubbleUp(i int) {
	p := bh.parent(i)
	for i > 0 && bh.a[i]-bh.a[p] > 0 {
		tmp := bh.a[i]
		bh.a[i] = bh.a[p]
		bh.a[p] = tmp

		i = p
		p = bh.parent(i)
	}
}

func (bh *BinaryHeap) Remove() int {
	x := bh.a[0]
	bh.n--
	bh.a[0] = bh.a[bh.n]
	bh.trickleDown(0)
	if 3*bh.n < len(bh.a) {
		bh.resize()
	}

	return x
}

func (bh *BinaryHeap) trickleDown(i int) {
	for {
		j := -1

		r := bh.right(i)
		if r < bh.n && bh.a[r]-bh.a[i] > 0 {
			l := bh.left(i)
			if bh.a[l]-bh.a[r] > 0 {
				j = l
			} else {
				j = r
			}
		} else {
			l := bh.left(i)
			if l < bh.n && bh.a[l]-bh.a[i] > 0 {
				j = l
			}
		}

		if j >= 0 {
			tmp := bh.a[i]
			bh.a[i] = bh.a[j]
			bh.a[j] = tmp

			i = j
		}

		if j < 0 {
			break
		}
	}
}

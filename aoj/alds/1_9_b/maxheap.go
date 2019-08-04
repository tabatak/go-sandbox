package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var h int

func main() {
	fmt.Scan(&h)

	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	a := make([]int, h+1)
	for i := 1; i < h+1; i++ {
		sc.Scan()
		a[i], _ = strconv.Atoi(sc.Text())
	}

	for i := h / 2; i >= 1; i-- {
		maxHeapify(a, i)
	}

	wr := bufio.NewWriter(os.Stdout)
	for i := 1; i < h+1; i++ {
		wr.Write([]byte(fmt.Sprintf(" %d", a[i])))
	}
	wr.Write([]byte("\n"))
	wr.Flush()
}

func maxHeapify(a []int, i int) {
	var l, r, largest int
	l = 2 * i
	r = 2*i + 1

	//左の子、自分、右の子で値が最大のノードを選ぶ
	if l <= h && a[l] > a[i] {
		largest = l
	} else {
		largest = i
	}
	if r <= h && a[r] > a[largest] {
		largest = r
	}

	if largest != i {
		tmp := a[largest]
		a[largest] = a[i]
		a[i] = tmp

		maxHeapify(a, largest)
	}
}

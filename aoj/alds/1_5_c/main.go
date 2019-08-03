package main

import (
	"fmt"
	"math"
)

var (
	n int
)

type Point struct {
	x float64
	y float64
}

func main() {
	fmt.Scan(&n)

	var a, b Point

	a.x = 0
	a.y = 0
	b.x = 100
	b.y = 0

	fmt.Printf("%.8f %.8f\n", a.x, a.y)
	koch(n, a, b)
	fmt.Printf("%.8f %.8f\n", b.x, b.y)

}

func koch(d int, a, b Point) {
	if d == 0 {
		return
	}

	var s, t, u Point
	th := math.Pi * 60.0 / 180.0

	s.x = (2.0*a.x + 1.0*b.x) / 3.0
	s.y = (2.0*a.y + 1.0*b.y) / 3.0
	t.x = (1.0*a.x + 2.0*b.x) / 3.0
	t.y = (1.0*a.y + 2.0*b.y) / 3.0
	u.x = (t.x-s.x)*math.Cos(th) - (t.y-s.y)*math.Sin(th) + s.x
	u.y = (t.x-s.x)*math.Sin(th) + (t.y-s.y)*math.Cos(th) + s.y

	koch(d-1, a, s)
	fmt.Printf("%.8f %.8f\n", s.x, s.y)
	koch(d-1, s, u)
	fmt.Printf("%.8f %.8f\n", u.x, u.y)
	koch(d-1, u, t)
	fmt.Printf("%.8f %.8f\n", t.x, t.y)
	koch(d-1, t, b)
}

package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	target := int64(3060521820)
	diff := math.Abs(float64(target - time.Now().Unix()))

	fmt.Printf("%f\n", diff)
	fmt.Println(diff > 31536000)
}

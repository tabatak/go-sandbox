package main

import (
	"fmt"
	"time"
)

func main() {
	var ym string
	fmt.Scan(&ym)

	t, _ := time.Parse("2006/01/02", ym+"/01")
	wd := int(t.Weekday())
	for i := 0; i <= wd; i++ {
		fmt.Print("   ")
	}
	tmp := t
	for {
		if tmp.Weekday() == 0 {
			fmt.Printf(" (%d) ", tmp.Day())
		} else if tmp.Weekday() == 6 {
			fmt.Printf(" [%d] ", tmp.Day())
		} else {
			fmt.Printf(" %2d ", tmp.Day())
		}
		wd++
		if wd > 6 {
			fmt.Println()
			wd = 0
		}

		tmp = tmp.AddDate(0, 0, 1)
		if tmp.Month() != t.Month() {
			break
		}
	}
}

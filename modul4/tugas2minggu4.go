package main

import (
	"fmt"
)

func hitung(s *int, t *int) {
	*t = 0
	*s = 0
	for i := 0; i < 8; i++ {
		var x int
		fmt.Scan(&x)
		if x <= 300 {
			*s++
			*t += x
		} else {
			*t += 301
		}
	}
}

func main() {
	var n string
	maxS, minT := -1, 1<<30
	var win string

	for {
		fmt.Scan(&n)
		if n == "Selesai" {
			break
		}

		var s, t int
		hitung(&s, &t)

		if s > maxS || (s == maxS && t < minT) {
			maxS = s
			minT = t
			win = n
		}
	}

	fmt.Println(win, maxS, minT)
}
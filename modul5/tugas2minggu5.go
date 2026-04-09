package main

import "fmt"

func pola(n, i int) {
	if i > n {
		return
	}

	for j := 1; j <= i; j++ {
		fmt.Print("*")
	}
	fmt.Println()

	pola(n, i+1)
}

func main() {
	var n int
	fmt.Scan(&n)

	pola(n, 1)
}
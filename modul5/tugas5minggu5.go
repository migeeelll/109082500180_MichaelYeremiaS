package main

import "fmt"

func ganjil(n int) {
	for i:= 1; i<= n;i++ {
		if i % 2 == 1{
			fmt.Printf("%d ",i)
		}
	}
}

func main() {
	var n int

	fmt.Scan(&n)

	ganjil(n)
}
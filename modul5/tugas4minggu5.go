package main

import "fmt"

func urut(n int) {
	for i:= n; i> 1;i-- {
		fmt.Printf("%d ",i)
		}
	for i:= 1; i<= n;i++ {
		fmt.Printf("%d ",i)
		}
	}
	



func main() {
	var n int

	fmt.Scan(&n)

	urut(n)
}
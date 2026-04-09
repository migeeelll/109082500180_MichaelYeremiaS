package main

import "fmt"

func pangkat(n,m int) {
	var x  int
	x=1
	for i:= 1; i<= m;i++ {
		x *=n
	}
	fmt.Print(x)
}

func main() {
	var n,m int

	fmt.Scan(&n,&m)

	pangkat(n,m)
}
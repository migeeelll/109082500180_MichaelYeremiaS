package main

import "fmt"

func f(n int) int {
	h := 1
	for i := 2; i <= n; i++ {
		h *= i
	}
	return h
}

func permutasi(n, r int) int {
	return f(n) / f(n-r)
}

func kombinasi(n, r int) int {
	return f(n) / (f(r) * f(n-r))
}

func main() {
	var a, b, c1, d int
	fmt.Scan(&a, &b, &c1, &d)

	fmt.Println(permutasi(a, c1), kombinasi(a, c1))
	fmt.Println(permutasi(b, d), kombinasi(b, d))
}
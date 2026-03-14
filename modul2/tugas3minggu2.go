package main

import "fmt"

func main() {
	var (
		a,b,c,d,e,f int64
		
	)
	fmt.Print("Masukan berat(g): ")
	fmt.Scan(&a)

	b = a / 1000
	c = a % 1000
	d = b * 10000
	e=0

	if c>=500{
			e = c * 5
		}else {
			e = c * 15
		}

	if b < 10{
		f = d + e 
	}else{
		f = d
	}

	
	fmt.Printf("Detail berat: %dkg + %dg\n",b,c)
	fmt.Printf("Detail biaya: Rp. %d + Rp. %d\n",d,e)
	fmt.Printf("Total biaya: Rp. %d\n",f)

	
	}

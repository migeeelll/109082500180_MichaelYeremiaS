package main

import "fmt"

func main() {
	var (
		a,b,c,d string
		h,h2 bool
	)
	h2 = true
	for i:=1;i<=5;i++{
		fmt.Printf("Percobaan %d : ",i)
		fmt.Scan(&a,&b,&c,&d)
		if a=="merah"&&b=="kuning"&&c=="hijau"&&d=="ungu"{
			h = true
		}else{
			h = false
		}
		h2 = h && h2
	}
	fmt.Printf("Hasil: %t",h2)
}

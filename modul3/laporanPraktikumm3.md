# <h1 align="center">Laporan Praktikum Modul 3 - fungsi minggu3 </h1>
<p align="center">Michael yeremia s - 109082500180</p>

## Unguided 

### 1. [Soal]
#### soal1.go

```go
package main

import "fmt"

func factorial(n int) int {
	hasil := 1
	for i := 1; i <= n; i++ {
		hasil *= i
	}
	return hasil
}

func permutation(n,r int) int{
	return factorial(n) / factorial(n-r)
}

func combination(n,r int) int{
	return factorial(n) / (factorial(r) * factorial(n-r))
}

func main(){
	var a,b,c,d int
	fmt.Scan(&a,&b,&c,&d)
	fmt.Println(permutation(a,c), combination(a,c))
	fmt.Println(permutation(b,d), combination(b,d))
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/migeeelll/109082500180_MichaelYeremiaS/blob/main/modul3/output/soal1.png)
[penjelasan]
code ini digunakan untuk menghitung permutasi dan kombinasi untuk bantuin jonas (walaupun sebenernya gak mau :p) inputan 4 bilangan asli keluaran nya hasil dari permutasi dan kombinasi

### 2. [Soal]
#### soal2.go

```go
package main

import "fmt"

func f(x int) int {
	return x * x
}

func g(x int) int {
	return x - 2
}

func h(x int) int {
	return x + 1
}

func main() {
	var a, b, c int

	fmt.Scan(&a, &b, &c)

	fmt.Println(f(g(h(a))))
	fmt.Println(g(h(f(b))))
	fmt.Println(h(f(g(c))))
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/migeeelll/109082500180_MichaelYeremiaS/blob/main/modul3/output/soal2.png)
[penjelasan]
code ini digunakan untuk menghitung 3 buah fungsi yg sudah tersedia. inpitan nya merupakan 3 buah bilangan bulat. keluaran merupakan hasil perkalian2 fungsi tersebut.

### 3. [Soal]
#### soal3.go

```go
package main

import (
	"fmt"
	"math"
)

func jarak(a, b, c, d float64) float64 {
	return math.Sqrt((a-c)*(a-c) + (b-d)*(b-d))
}

func didalam(cx, cy, r, x, y float64) bool {
	return jarak(cx, cy, x, y) <= r
}

func main() {
	var cx1, cy1, r1 float64
	var cx2, cy2, r2 float64
	var x, y float64

	fmt.Scan(&cx1, &cy1, &r1,&cx2, &cy2, &r2,&x, &y)

	d1 := didalam(cx1, cy1, r1, x, y)
	d2 := didalam(cx2, cy2, r2, x, y)

	if d1 && d2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if d1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if d2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/migeeelll/109082500180_MichaelYeremiaS/blob/main/modul3/output/soal2.png)
[penjelasan]
sama kek di soal

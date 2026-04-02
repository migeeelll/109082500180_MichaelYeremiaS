# <h1 align="center">Laporan Praktikum Modul 4 - PROSEDUR </h1>
<p align="center">Michael yeremia s - 109082500180</p>

## Unguided 

### 1. [Soal]
#### tugas1minggu2.go

```go
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

```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/migeeelll/109082500180_MichaelYeremiaS/blob/main/modul4/output/image.png)
[penjelasan]
ya gitu deh liat soalnya aja kak

### 2. [Soal]
#### tugas2minggu2.go

```go
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

```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/migeeelll/109082500180_MichaelYeremiaS/blob/main/modul4/output/image%20copy.png)
[penjelasan]
buat ngitung scor(kata soal)

### 3. [Soal]
#### tugas3minggu2.go

```go
package main

import "fmt"

func cetakDeret(n int) {
	for n != 1 {
		fmt.Print(n, " ")
		if n%2 == 0 {
			n = n / 2
		} else {
			n = 3*n + 1
		}
	}
	fmt.Print(1)
}

func main() {
	var n int

	fmt.Scan(&n)

	cetakDeret(n)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/migeeelll/109082500180_MichaelYeremiaS/blob/main/modul4/output/image%20copy%202.png)
[penjelasan]
deret bilangan kalo genap bagi dua,kalo ganjil kali 3 tambah 1 (coba liat soalnya)
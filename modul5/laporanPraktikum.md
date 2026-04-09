# <h1 align="center">Laporan Praktikum Modul 4 - PROSEDUR </h1>
<p align="center">Michael yeremia s - 109082500180</p>

## Unguided 

### 1. [Soal]
#### tugas1minggu2.go

```go
package main

import "fmt"

func cetakNFibo(n int) {
	var f1, f2, f3 int
	f2 = 0
	f3 = 1
	for i := 1; i <= n; i++ {
		fmt.Println(f3)
		f1 = f2
		f2 = f3
		f3 = f1 + f2
	}
}

func main() {
	var n int

	fmt.Scan(&n)

	cetakNFibo(n)
}

```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/migeeelll/109082500180_MichaelYeremiaS/blob/main/modul5/output/Screenshot%202026-04-10%20001652.png)
[penjelasan]
ya gitu deh liat soalnya aja kak (Deret fibonacci adalah sebuah deret dengan nilai suku ke-0 dan ke-1 adalah 0 dan 1, dan nilai
suku ke-n selanjutnya adalah hasil penjumlahan dua suku sebelumnya. Secara umum dapat
diformulasikan Sn = Sn−1 + Sn−2 . Berikut ini adalah contoh nilai deret fibonacci hingga suku
ke-10. Buatlah program yang mengimplementasikan fungsi rekursif pada deret fibonacci
tersebut.)

### 2. [Soal]
#### tugas2minggu2.go

```go
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
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/migeeelll/109082500180_MichaelYeremiaS/blob/main/modul5/output/Screenshot%202026-04-10%20001718.png)
[penjelasan]
sebuah program yang digunakan untuk menampilkan pola bintang berikut ini dengan menggunakan fungsi rekursif. N adalah masukan dari user.

### 3. [Soal]
#### tugas3minggu2.go

```go
package main

import "fmt"

func faktor(n int) {
	for i:= 1; i<= n;i++ {
		if n % i == 0{
			fmt.Println(i)
		}
	}
}

func main() {
	var n int

	fmt.Scan(&n)

	faktor(n)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/migeeelll/109082500180_MichaelYeremiaS/blob/main/modul5/output/Screenshot%202026-04-10%20001932.png)
[penjelasan]
program yang mengimplementasikan rekursif untuk menampilkan faktor bilangan dari suatu N, atau bilangan yang apa saja yang habis membagi N.

### 4. [Soal]
#### tugas3minggu2.go

```go
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
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/migeeelll/109082500180_MichaelYeremiaS/blob/main/modul5/output/Screenshot%202026-04-10%20001953.png)
[penjelasan]
program yang mengimplementasikan rekursif untuk menampilkan barisan bilangan tertentu.

### 5. [Soal]
#### tugas3minggu2.go

```go
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
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/migeeelll/109082500180_MichaelYeremiaS/blob/main/modul5/output/Screenshot%202026-04-10%20002017.png)
[penjelasan]
program yang mengimplementasikan rekursif untuk menampilkan barisan bilangan ganjil.

### 6. [Soal]
#### tugas3minggu2.go

```go
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
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/migeeelll/109082500180_MichaelYeremiaS/blob/main/modul5/output/Screenshot%202026-04-10%20002036.png)
[penjelasan]
program yang mengimplementasikan rekursif untuk mencari hasil pangkat dari dua buah bilangan.

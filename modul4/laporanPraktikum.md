# <h1 align="center">Laporan Praktikum Modul 2 - REVIEW ALGORITMA DAN PEMROGRAMAN 1 </h1>
<p align="center">Michael yeremia s - 109082500180</p>

## Unguided 

### 1. [Soal]
#### tugas1minggu2.go

```go
package main

import "fmt"

func main() {
	var (
		satu, dua, tiga string
		temp            string
	)
	fmt.Print("Masukan input string: ")
	fmt.Scanln(&satu)
	fmt.Print("Masukan input string: ")
	fmt.Scanln(&dua)
	fmt.Print("Masukan input string: ")
	fmt.Scanln(&tiga)
	fmt.Println("Output awal = " + satu + " " + dua + " " + tiga)
	temp = satu
	satu = dua
	dua = tiga
	tiga = temp
	fmt.Println("Output akhir = " + satu + " " + dua + " " + tiga)
}

```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/shellyneu/103112430114_Shellyn/blob/main/modul1/output/output-soal1.png)
[penjelasan]
code ini memindah kan 

### 2. [Soal]
#### tugas2minggu2.go

```go
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


```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/shellyneu/103112430114_Shellyn/blob/main/modul1/output/output-soal1.png)
[penjelasan]

### 3. [Soal]
#### tugas3minggu2.go

```go
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


```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/shellyneu/103112430114_Shellyn/blob/main/modul1/output/output-soal1.png)
[penjelasan]

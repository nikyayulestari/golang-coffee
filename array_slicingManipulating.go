package main 

import (  
    "fmt"
)

func main() {  
	x := [5]int {1, 2, 3, 4, 5}
	fmt.Println("Array awal", x)
	
	dislice := x[2:4] // akan memotong array ke 2 sampai 4 (ingat, 4 tidak termasuk), jadi angka 3 dan 4
	fmt.Println("Yang dipotong", dislice)
	
	for i := range dislice {
		dislice[i]++ // mengganti nilai yang indeksnya ada di dislice dengan indeks setelahnya (++1) sampai indeks tersebut bukan lagi indeks yang merupakan indeks pada dislice.
	}
	
	fmt.Println("Hasil potongan", x)
	
	// Hasilnya [1 2 4 5 5], 3 dan 4 berubah jadi 4 dan 5 karena di +1 pada perulangan fornya.
}

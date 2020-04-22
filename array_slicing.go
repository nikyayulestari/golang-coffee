package main 

import (  
    "fmt"
)

func main() {  
	x := [5]int {1, 2, 3, 4, 5}
	y := x[1:4] // akan memanggil array, mulai dari array ke 1 sampai dengan batasnya array ke 4, array indeks ke 4 tidak ditulis.
	fmt.Println(y)
	y = x[2:] // akan memanggil array, mulai dari array ke 2 sampai akhir.
	fmt.Println(y)
	y = x[:3] // akan memanggil array, mulai dari array awal hingga array ke 3, array indeks 3nya tidak masuk.
	fmt.Println(y)
}

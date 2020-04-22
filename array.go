package main 

import (  
    "fmt"
)

func main() {  
	// how to declare and define array #1
	var a [2]int
	a[0] = 23
	a[1] = 10
	fmt.Println(a)
	
	// how to declare and define array #2
	b := [3]string{"aku","kamu","kita"}
	fmt.Println(b)
	
	// dari 3 kapasitas array yang dibuat, yang diisi hanya array pertama atau array indeks 0, sehingga 2 array lainnya akan bernilai 0
	c := [3]int{12}
	fmt.Println(c)
	
	// begitupun apabila dibuat array tanpa menginisiasi nilainya, maka semua array akan bernilai 0
	var d [3]int
	fmt.Println(d)
}

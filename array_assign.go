package main 

import (  
    "fmt"
)

func main() {  
	// array a tidak diberi batasan nilai [...] oleh karena itu dapat diinisiasi dengan nilai sampai berapapun.
	a := [...]string{"USA", "China", "India", "Germany", "France"}
	
	b := a
	// nilai b berisikan nilai pada array a, bisa, karena a tidak ditentukan panjang arraynya, begitupun b.
	//Pada array, meskipun sama" bentuk array tetapi apabila panjangnya berbeda, misal a[3] dan b[5] maka tetap tidak bisa dioperasikan.
	
	
	b[0] = "Singapore" // mengganti indeks 0 pada b menjadi Singapore, artinya replace yang tadinya USA.
	
	fmt.Println("a is", a)
	fmt.Println("b is", b) 
}

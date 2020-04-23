package main  
import (  
	"fmt"
	"sync" // import sync untuk menggunakan Mutex & WaitGroup
)

var x  = 0  
func increment(wg *sync.WaitGroup, m *sync.Mutex) { // buat parameter dengan tipe WaitGroup dan Mutex, penulisannya pake sync.
	m.Lock() // buka dulu mutexnya dengan lock untuk menghindari tabrakan dan program fokus menjalankan satu operasi saja.
	x = x + 1  // ini operasi yang dijalankan.
	m.Unlock() // setelah operasi selesai maka m dengan tipe mutex tadi dibuka/di unlock, sehingga membuka kesempatan bagi program untuk menjalankan perintah lainnya.
	wg.Done()   // apabila operasi yang dibutuhkan pada fungsi ini telah dijalankan semua, maka wg dengan tipe WaitGroup tadi dapat diselesaikan (.Done()).
}

func main() {  
	// dibuat variable yang menggantikan kedua parameter diatas wg (dengan tipe WaitGroup) dan m (Mutex) yaitu w dan m.
	var w sync.WaitGroup
	var m sync.Mutex
	
	for i := 0; i < 1000; i++ { // perulangan dilakukan sebanyak 1000 kali.
		w.Add(1) // Add guna menginisiasi value yang dibawa oleh var w adalah 1.
		go increment(&w, &m) // goroutines yang memanggil fungsi increment diatas dengan parameter w dan m.
	}
	w.Wait() // method wait() digunakan untuk menunggu dan memastikan bahwa operasi sebelumnya (perulangan) telah benar" selesai dilakukan sebelum dilanjutkan pada baris perintah selanjutnya, yaitu line 25.
	fmt.Println("final value of x", x) // mengeluarkan hasil dari penjumlahan pada looping diatas, yakni 1000.
}

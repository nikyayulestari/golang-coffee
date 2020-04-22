package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s { // v = value, for _ range s artinya memanggil semua nilai dari variabel s secara berurutan, kemudian menyimpannya pada variabel v untuk func ini.
		sum += v // menjumlahkan semua nilai yang dipanggil sesuai parameter s nantinya.
	}
	c <- sum // menyimpan nilai hasil sum ke parameter c (tipenya chan int).
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	//  channel harus di "make" atau dibuat dl sbelum dipake (pd fungsi).
	
	go sum(s[:len(s)/2], c)
	// mengambil array s mulai indeks awal hingga ke-3 (len(s) itu membaca jumlah karakter/isi dari var s = 6, disitu len(s) kan 6, jadi len(s)/2 = 3).
	
	go sum(s[len(s)/2:], c)
	// sama kayak diatas berarti mengambil dari array indeks ke-3 sampai akhir.
	
	x, y := <-c, <-c
	// x dan y receive from c, anehnya diterimanya kebalik, jadi goroutines yang paling akhir itu yang akan dieksekusi dluan (disimpan pd x), dan satunya pd y.

	fmt.Println(x, y, x+y)
}

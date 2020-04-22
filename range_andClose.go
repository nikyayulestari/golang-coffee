package main

import (
	"fmt"
)

func fibonacci(n int, c chan int) {
	x, y := 0, 1 // inisiasi x = 0 dan y = 1
	for i := 0; i < n; i++ { // i akan berulang sampai 10kali (n)
		c <- x
		x, y = y, x+y
	}
	close(c) // close the channel
}

func main() {
	c := make(chan int, 10) // membuat c sebanyak 10 hasil
	go fibonacci(cap(c), c) // cap(c) itu untuk mengambil 10 nya itu.
	for i := range c { // diulang range (sampai dengan) c nya abis, alias 10 kali tadi sesuai dengan jumlah yang dideklarasikan.
		fmt.Println(i)
	}
}

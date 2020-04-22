package main

import "fmt"

func sendData(sendch chan<- int) {
	sendch <- 10 // mendefinisikan bahwa channel sendch hanya akan menjadi fungsi pengirim, yaitu mengirimkan angka 10
}

func main() {
	chnl := make(chan int)
	go sendData(chnl)
	fmt.Println(<-chnl) // mengeluarkan hasil dari fungsi sendData dengan <-chnl.
}

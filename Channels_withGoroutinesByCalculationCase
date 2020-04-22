package main

import (
	"fmt"
)

func calcSquares(number int, squareop chan int) {
	sum := 0
	for number != 0 {
		digit := number % 10
		sum += digit * digit
		number /= 10
	}
	squareop <- sum
}

func calcCubes(number int, cubeop chan int) {
	sum := 0
	for number != 0 {
		digit := number % 10
		sum += digit * digit * digit
		number /= 10
	}
	cubeop <- sum
}

func main() {
	number := 589 // angka yang bakal dipake perhitungan (untuk ngisi parameter number di masing" fungsi)
	sqrch := make(chan int) // deklarasi var untuk simpan hasil fungsi chan intnya (calcSquares yaitu squareop)
	cubech := make(chan int) // deklarasi var untuk simpan hasil fungsi chan intnya (calcCubes yaitu cubeop)
	go calcSquares(number, sqrch) // dibuat goroutines untuk masing" fungsi
	go calcCubes(number, cubech)
	squares, cubes := <-sqrch, <-cubech // cara inisiasi langsung 2 channels
	fmt.Println("Final output", squares+cubes)
}

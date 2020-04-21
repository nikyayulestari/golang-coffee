package main

import "fmt"

func main() {
	var a chan int // deklarasi a (null)
	if a == nil { // null in go is "nil"
		fmt.Println("channel a is nil, going to define it")
		a = make(chan int) // make use to define the channel (write by chan)
		fmt.Printf("Type of a is %T", a) // print out (%) type of (T is for print of type-%T) variable a
	}
}

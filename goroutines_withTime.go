package main

import (
	"fmt"
	"time"
)

func hello() {
	fmt.Println("Hello world goroutine")
}
func main() {
	go hello() //eksekusi fungsi hello()
	time.Sleep(3 * time.Second) //waiting 3 seconds to run next line (line 15)
	fmt.Println("main function")
}

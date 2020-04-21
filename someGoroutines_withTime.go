package main

import (
    "fmt"
    "time"
)

func numbers() {  
    for i := 1; i <= 5; i++ {
        time.Sleep(250 * time.Millisecond) // waiting time
        fmt.Printf("%d ", i)
    }
}
func alphabets() {  
    for i := 'a'; i <= 'e'; i++ {
        time.Sleep(400 * time.Millisecond) // waiting time
        fmt.Printf("%c ", i)
    }
}
func main() {
    // print fungsi numbers() dan alphabets() secara bergantian
    go numbers() // goroutines function numbers()
    go alphabets() // goroutines function alphabet()
    time.Sleep(3000 * time.Millisecond) // waiting time
    fmt.Println("main terminated")
}

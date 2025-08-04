package main

import (
	"fmt"
	"time"
)

func display(s string) {
	for i := 0; i < 3; i++ {
		time.Sleep(500 * time.Millisecond)
		fmt.Println(s)
	}
}
func printNumber() {
	i := 0
	for {
		time.Sleep(1 * time.Second)
		i++
		fmt.Println(i)
	}
}
func main() {
	fmt.Println("launch goroutine")
	go printNumber()
	fmt.Println("launch goroutine")
	go printNumber()
	time.Sleep(1 * time.Minute)
}

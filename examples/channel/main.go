package main

import (
	"fmt"
	"time"
)

func send(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
		time.Sleep(time.Second * 1)
	}
	close(ch)
}

func recv(ch chan int) {
	time.Sleep(time.Second * 1)
	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}
}

func main() {
	fmt.Println("start channel...")
	ch := make(chan int, 10)
	// defer close(ch)
	go send(ch)
	go recv(ch)
	time.Sleep(time.Second * 30)
}

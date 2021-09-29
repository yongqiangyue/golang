package main

import (
	"fmt"
	"time"
)

// note: go build -o busyloop main.go
func main() {
	go func() {
		for {
			fmt.Println("log from go")
			time.Sleep(time.Duration(2) * time.Second)
		}
	}()
	for {
		fmt.Println("main from go")
		time.Sleep(time.Duration(1) * time.Second)
	}
}

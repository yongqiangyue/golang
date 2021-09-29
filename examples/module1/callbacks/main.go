package main

import (
	"fmt"
)

func main() {
	// var a *int
	// *a += 1
	fmt.Println(DoOperation(1, increase))
	fmt.Println(DoOperation(1, decrease))
}

func DoOperation(y int, f func(int, int) int) int {
	return f(y, 1)
}

func increase(a, b int) int {
	return a + b
}

func decrease(a, b int) int {
	return a - b
}

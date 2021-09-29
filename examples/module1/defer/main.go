package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// defer fmt.Println("1")
	// defer fmt.Println("2")
	// defer fmt.Println("3")
	loopFunc()
	time.Sleep(time.Second)
	// F()
}

func G() {
	defer func() {
		fmt.Println("c")
	}()
	F()
	fmt.Println("继续执行")
}

func F() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("捕获异常:", err)
		}
		fmt.Println("b")
	}()
	panic("a")
}

func loopFunc() {
	lock := sync.Mutex{}
	for i := 0; i < 3; i++ {
		// go unc(i int) {
		func(i int) {
			lock.Lock()
			defer lock.Unlock()
			fmt.Println("loopFunc:", i)
		}(i)
		// }(i)
	}
}

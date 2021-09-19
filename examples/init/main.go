package main

// 什么情况下，只引用一个包的init函数
import (
	"fmt"

	_ "github.com/cncamp/golang/examples/init/a"
	b "github.com/cncamp/golang/examples/init/b"
)

func init() {
	fmt.Println("main init")
	//b.MyVal = "main"
}
func main() {
	fmt.Println(b.MyVal)
}

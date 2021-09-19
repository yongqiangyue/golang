package b

import (
	"fmt"
)

var MyVal string = "xx"

func init() {
	fmt.Println("init from b")
	MyVal  = "b"
}

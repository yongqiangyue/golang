package a

import (
	"fmt"

	b "github.com/cncamp/golang/examples/init/b"
)

func init() {
	fmt.Println("init from a")
	b.MyVal = "a"
}

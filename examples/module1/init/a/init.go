package a

import (
	"fmt"

<<<<<<< HEAD:examples/init/a/init.go
	b "github.com/cncamp/golang/examples/init/b"
=======
	_ "github.com/cncamp/golang/examples/module1/init/b"
>>>>>>> upstream/master:examples/module1/init/a/init.go
)

func init() {
	fmt.Println("init from a")
	b.MyVal = "a"
}

package main

import (
	"flag"
	"fmt"
)

/* note:
go run main.go --name ft
go run main.go --name ai.ft
*/
func main() {
	name := flag.String("name", "world", "specify the name you want to say hi")
	flag.Parse()
	// var strVal string
	// flag.StringVar(&strVal, "sval", "zhangshicun", "notes")
	// flag.Parse()
	// fmt.Println("sval:", strVal)
	// fmt.Println("os args is:", os.Args)
	// fmt.Println("input parameter is:", *name)
	// fullString := fmt.Sprintf("Hello %s from Go\n", *name)
	// fmt.Println(fullString)
	result, err := duplicateString(*name)
	if err == nil {
		fmt.Println(result)
	} else {
		fmt.Println(err)
	}
}

func duplicateString(input string) (result string, err error) {
	if input == "ft" {
		err = fmt.Errorf("ft is not allowed")
	} else {
		result = input + input
	}
	return
}

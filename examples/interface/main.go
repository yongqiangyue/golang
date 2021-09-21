package main

import (
	"encoding/json"
	"fmt"
)

type IF interface {
	getName() string
}

type Human struct {
	FirstName, LastName string
	Age                 int
}

// 字符串转struct
func unmarshal2Struct(humanStr string) *Human {
	h := Human{}
	if err := json.Unmarshal([]byte(humanStr), &h); err != nil {
		fmt.Println("json unmarshal failed.", err)
	}
	fmt.Printf("call %p\n", &h)
	defer fmt.Printf("defer call %p\n", &h)
	return &h
}

// struct转字符串
func marshal2JsonString(h Human) string {
	h.Age = 37
	updatedBytes, err := json.Marshal(&h)
	if err != nil {
		fmt.Println(err)
	}
	return string(updatedBytes)
}

type Plane struct {
	vendor string
	model  string
}

func (h *Human) getName() string {
	return h.FirstName + "," + h.LastName
}

func (p Plane) getName() string {
	return fmt.Sprintf("vendor: %s, model: %s", p.vendor, p.model)
}

type Car struct {
	factory, model string
}

func (c *Car) getName() string {
	return c.factory + "-" + c.model
}

func main() {
	h := Human{FirstName: "yue", LastName: "yongqianng", Age: 30}
	// fmt.Println(h)
	hString := marshal2JsonString(h)
	// fmt.Println(hString)
	// fmt.Println(h)

	hh := unmarshal2Struct(hString)
	fmt.Printf("%p\n", hh)
	// interfaces := []IF{}
	// h := new(Human)
	// h.firstName = "first"
	// h.lastName = "last"
	// interfaces = append(interfaces, h)
	// c := new(Car)
	// c.factory = "benz"
	// c.model = "s"
	// interfaces = append(interfaces, c)
	// for _, f := range interfaces {
	// 	fmt.Println(f.getName())
	// }
	// p := Plane{}
	// p.vendor = "testVendor"
	// p.model = "testModel"
	// fmt.Println(p.getName())
}

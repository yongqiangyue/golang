package main

import (
	"fmt"
	"reflect"
)

type MyType struct {
	Name string `json:"name"`
}

type rect_shape struct {
	x float32
	y float32
}

func (shape rect_shape) erea() float32 {
	return shape.x * shape.y
}

type myInt int

func (i *myInt) addNumber(n int) int {
	return n + n
}

type mySlice []int

func (v mySlice) sumOfSlice() int {
	sum := 0
	for index, _ := range v {
		sum += v[index]
	}
	return sum
}

// note: 方法：在go中和struct关联，被称为struct的receiver.
func main() {
	x := mySlice{1, 2, 3, 4, 5}
	fmt.Println(x.sumOfSlice())

	var i *myInt = new(myInt)
	fmt.Println(i.addNumber(4))

	var shape rect_shape = rect_shape{400, 200}
	fmt.Println(shape.erea())

	mt := MyType{Name: "test"}
	myType := reflect.TypeOf(mt)
	name := myType.Field(0)
	tag := name.Tag.Get("json")
	println(tag)
	tb := TypeB{P2: "p2", TypeA: TypeA{P1: "p1"}}
	//可以直接访问 TypeA.P1
	println(tb.P1)
}

type TypeA struct {
	P1 string
}

type TypeB struct {
	P2 string
	TypeA
}

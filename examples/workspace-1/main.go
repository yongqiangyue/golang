package main

import (
	"fmt"
)

/* 给定一个字符串数组
["I","am","stupid","and","weak"]
用 for 循环遍历该数组并修改为
["I","am","smart","and","strong"]*/

func main() {
	fmt.Println("workspace-1 start")
	myArray := [...]string{"I", "am", "stupid", "and", "weak"}
	mySlice := myArray[:]
	// fmt.Printf("myArray：%+v\n", myArray)
	fmt.Println("myArray:", myArray)
	for i, _ := range mySlice {
		if i == 2 {
			mySlice[i] = "smart"
		}
		if i == 4 {
			mySlice[i] = "strong"
		}
	}
	fmt.Println("myArray:", myArray)
	// fmt.Printf("myArray：%+v\n", myArray)

}

package main

import (
	"fmt"
)

func main() {
	num1 := 100
	fmt.Printf("num1的类型是:%T  num1的值是:%v  num1的地址是:%v\n", num1, num1, &num1)

	num2 := new(int)
	// num2的类型是 => *int
	// num2的值是 => 指向0的地址
	// num2的地址
	fmt.Printf("num2的类型是:%T  num2的值是:%v  num2的地址是:%v\n", num2, num2, &num2)
	fmt.Println(*num2)
}

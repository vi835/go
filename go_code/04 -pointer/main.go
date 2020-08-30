package main

import (
	"fmt"
)

// 演示golang中的指针类型
func main() {

	// 基本数据类型在内存布局
	var i int = 10
	// i的地址是什么 &i
	fmt.Println("i的地址====>", &i)

	// 下面的 var ptr *int = &i
	// 1 ptr是一个指针变量
	// 2 ptr的类型是*int
	// 3 ptr本身的值是 &i
	var ptr *int = &i
	fmt.Printf("ptr======> %v\n", ptr)
	fmt.Printf("ptr 的地址是======> %v\n", &ptr)
	fmt.Printf("ptr 指向的值是=====> %v\n", *ptr)
}

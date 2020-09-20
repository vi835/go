package main

import "fmt"

func main() {
	// 切片的拷贝操作
	// 切片使用copy内置函数完成拷贝 它们之间的数据空间是独立的
	var slice1 []int = []int{1, 2, 3, 4, 5}
	var slice2 []int = make([]int, 10)
	copy(slice2, slice1)
	fmt.Println("slice1=", slice1)
	fmt.Println("slice2=", slice2)
}

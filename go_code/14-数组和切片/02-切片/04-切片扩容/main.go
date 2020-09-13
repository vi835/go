package main

import "fmt"

func main() {
	// 用append内置函数 可以对切片进行动态追加
	var slice []int = []int{100, 200, 300}
	fmt.Println("slice=", slice)

	// 通过append直接给slice追加具体的元素
	slice = append(slice, 100)
	fmt.Println("slice=", slice)

	slice = append(slice, 1, 2)
	fmt.Println("slice=", slice)

	// 追加一个切片 需要使用...
	slice = append(slice, slice...)
	fmt.Println("slice=", slice)
}

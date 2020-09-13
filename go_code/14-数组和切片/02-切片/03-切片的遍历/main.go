package main

import "fmt"

func main() {
	var arr [5]int = [...]int{10, 20, 30, 40, 50}
	slice := arr[1:4] // 20 20,30,40
	// 使用常规的for循环遍历切片
	for i := 0; i < len(slice); i++ {
		fmt.Printf("slice[%v]=%v\n", i, slice[i])
	}
	fmt.Println()
	// 使用for-range遍历
	for i, v := range slice {
		fmt.Printf("slice[%v]=%v\n", i, v)
	}

	// 切片还可以继续切片
	slice2 := slice[1:2]
	// 因为arr,slice,slice2指向的数据空间是同一个 因为slice2[0]改变了
	// arr slice也都变化了
	slice2[0] = 100
	fmt.Println(slice2)
}

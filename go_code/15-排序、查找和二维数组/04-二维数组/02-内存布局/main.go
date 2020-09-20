package main

import "fmt"

func main() {
	// 以这个为例 来分析arr在内存中的布局
	var arr [2][2]int
	arr[1][1] = 10
	fmt.Println("arr=", arr)

	fmt.Printf("arr[0]的地址为%p\n", &arr[0])
	fmt.Printf("arr[1]的地址为%p\n", &arr[1])

	fmt.Printf("arr[0][0]的地址为%p\n", &arr[0][0])
	fmt.Printf("arr[1][0]的地址为%p\n", &arr[1][0])
}

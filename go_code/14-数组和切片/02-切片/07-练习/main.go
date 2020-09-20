package main

import "fmt"

// 编写一个函数 要求 可以接收一个 n int
// 能够将斐波那契数列放到切片中
func test(n int) []uint64 {
	// 声明一个切片 切片大小 n
	slice := make([]uint64, n)
	for i := 0; i < n; i++ {
		if i == 0 || i == 1 {
			slice[i] = 1
		} else {
			slice[i] = slice[i-1] + slice[i-2]
		}
	}
	return slice
}
func main() {
	fmt.Println(test(3))
}

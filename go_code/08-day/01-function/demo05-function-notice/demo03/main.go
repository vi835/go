package main

import (
	"fmt"
)

// 编写一个函数 可以求出1到多个int的值
// 可变参数的使用
func sum(n1 int, args ...int) int {
	sum := n1
	// 遍历args
	for i := 0; i < len(args); i++ {
		sum += args[i] // args[0] 表示取出args切片的第一个元素 其他一次类推
	}
	return sum
}
func main() {
	// 测试一下可变参数的使用
	res4 := sum(10)
	fmt.Println("res4=", res4)

	res5 := sum(10,1,-1,100,20)
	fmt.Println("res5=", res5)
}

package main

import "fmt"

func main() {
	// 在定义匿名函数时 就直接调用 这种匿名函数只能调用一次
	// 使用匿名函数的方式 求两个数的和
	res1 := func(n1 int, n2 int) int {
		return n1 + n2
	}(3, 4)
	fmt.Println("res1 = ", res1)

	// 将匿名函数付给一个变量(函数变量) 在通过该变量来调用匿名函数
	// 则a的数据类型就是函数类型 此时 我们可以通过a来完成调用
	a := func(n1 int, n2 int) int {
		return n1 - n2
	}
	res2 := a(10, 20)
	fmt.Println("res2 = ", res2)
}

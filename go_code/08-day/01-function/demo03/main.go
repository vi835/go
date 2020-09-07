package main

import (
	"fmt"
)

func test(n1 int) {
	n1 = n1 + 1
	fmt.Println("test() n1 =>", n1)
}

// 计算两个数的和
func getSum(n1 int, n2 int) int {
	sum := n1 + n2
	fmt.Println("func getSum = ", sum)
	// 当函数有return语句时 就是将结果返回给调用者
	return sum
}

// 编写一个函数 可以计算两个数的和和差 并返回结果
func getSumAndSub(n1 int, n2 int) (int, int) {
	sum := n1 + n2
	sub := n1 - n2
	return sum, sub
}

func main() {
	n1 := 10
	//调用test
	test(n1)
	fmt.Println("main() n1 =>", n1)

	sum := getSum(10, 29)
	fmt.Println("main sum =>", sum)

	// 调用getSumAndSub
	res1, res2 := getSumAndSub(10, 4)
	fmt.Println("res1: ", res1, " res2:", res2)

	// 希望忽略某个返回值 则使用 _ 符号表示占位符
	_,res3 := getSumAndSub(1, 3)
	fmt.Println("res3: ", res3)
}

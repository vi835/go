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

func main() {
	n1 := 10
	//调用test
	test(n1)
	fmt.Println("main() n1 =>", n1)

	sum :=getSum(10,29)
	fmt.Println("main sum =>",sum)
}

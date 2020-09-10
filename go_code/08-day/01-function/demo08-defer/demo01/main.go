package main

import "fmt"

func sum(n1 int, n2 int) int {
	// 当执行到defer时 暂时不执行 会将defer后面的语句压入到独立的栈中
	// 当函数执行完毕后 再从defer栈 按照陷入后出的方式出栈 执行
	defer fmt.Println("ok1 n1=", n1) //  order 3
	defer fmt.Println("ok2 n2=", n2) //  order 2
	res := n1 + n2
	fmt.Println("ok3 res=", res) // order 1
	return res
}

func main() {
	res := sum(10, 20)
	fmt.Println("res=", res) // order 4
}

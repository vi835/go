package main

import "fmt"

// 函数中的变量是局部的 函数外不生效
func test() {
	// n1 是test函数的局部变量 只能在test函数中使用
	var n1 = 10
	fmt.Println("n1:", n1)
}

func test02(n1 int) {
	n1 = n1 + 10
	fmt.Println("test02() n1 = ", n1)
}

func test03(n1 *int) {
	*n1 = *n1 + 10
	fmt.Println("test02() &n1 = ", *n1)
	fmt.Println("test02() &n1 = ", &n1)
}
func main() {
	// 这里不能使用n1 因为n1是test函数的局部变量
	// fmt.Println("n1:", n1)
	n1 := 20
	test02(n1)
	fmt.Println("main() n1 = ", n1)
	fmt.Println("--------------")
	n2 := 20
	test03(&n2)
	fmt.Println("main() n2 = ", n2)
}

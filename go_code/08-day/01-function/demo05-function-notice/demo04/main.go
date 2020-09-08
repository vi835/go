package main

import "fmt"

// 请编写一个函数swap(n1 *int,n2 *int) 可以交换n1和n2的值
func swap(n1 *int, n2 *int) {
	var temp int = *n1
	*n1 = *n2
	*n2 = temp
}
func main() {
	a := 10
	b := 20
	swap(&a, &b)
	fmt.Printf("a=%d,b=%d", a, b)
}

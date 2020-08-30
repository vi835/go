package main

import (
	"fmt"
)

// 演示Golang中字符类型使用
func main() {
	var c1 byte = 'a'
	var c2 byte = '0' // 字符的0

	// 当我们直接输出byte值时 就是输出了对应的字符的码值
	fmt.Println("c1 =", c1)
	fmt.Println("c2 =", c2)

	// 如果我们希望输出对应的字符 需要使用格式化输出
	fmt.Printf("c1 =%c c2 = %c\n", c1, c2)

	// var c3 byte = '北' // overflow溢出
	var c3 int = '北'
	fmt.Println("c3 =", c3)
	fmt.Printf("c3 = %c\n", c3)

	// 可以直接给某个变量赋一个数字 然后按格式化输出%c 会输出该数字对应的unicode字符
	var c4 int = 22269 // 22269 <=>国
	fmt.Printf("c4 = %c\n", c4)

	// 字符类型是可以进行运算的 相当于一个整数 因为他都有对应的Unicode码
	var n1 = 10 + 'a' // 10+97 =107
	fmt.Println("n1 =", n1)
	fmt.Printf("n1 = %c\n", n1)

}

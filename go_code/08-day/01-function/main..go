package main

import (
	"fmt"
)

func main() {
	// 输入两个数 在输入一个运算符(+,-,*,/) 得到结果
	var a float64
	var b float64
	var mark byte
	var result float64
	fmt.Println("请输入两个数 和一个运算符:")
	fmt.Scanf("%f %f %c", &a, &b, &mark)
	switch mark {
	case '+':
		result = a + b
	case '-':
		result = a - b
	case '*':
		result = a * b
	case '/':
		result = a / b
	default:
		fmt.Println("请输入正确的符号")
	}
	fmt.Println(result)
}

package main

import (
	"fmt"
)

// 将计算的功能 放到一个函数中 然后在需要使用 调用即可
func cal(a float64, b float64, mark byte) float64 {
	var result float64
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
	return result
}

func main() {
	// 输入两个数 在输入一个运算符(+,-,*,/) 得到结果
	var a float64
	var b float64
	var mark byte
	fmt.Println("请输入两个数 和一个运算符:")
	fmt.Scanf("%f %f %c", &a, &b, &mark)
	result :=cal(a,b,mark)
	fmt.Println(result)
}

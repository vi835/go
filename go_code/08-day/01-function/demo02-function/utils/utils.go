package utils

import (
	"fmt"
)

// 将计算的功能 放到一个函数中 然后在需要使用 调用即可
// 为了将其他包文件可以使用Cal函数 需要将c大写 类似其他编程语言的public
// go中该函数可导出
func Cal(a float64, b float64, mark byte) float64 {
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

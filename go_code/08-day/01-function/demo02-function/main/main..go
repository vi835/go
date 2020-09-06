package main

import (
	"fmt"
	"go_code/08-day/01-function/demo02-function/utils"
)

func main() {
	// 输入两个数 在输入一个运算符(+,-,*,/) 得到结果
	var a float64
	var b float64
	var mark byte
	var result float64
	fmt.Println("请输入两个数 和一个运算符:")
	fmt.Scanf("%f %f %c", &a, &b, &mark)
	result = utils.Cal(a, b, mark)
	fmt.Println(result)
}

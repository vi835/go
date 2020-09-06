package main

import (
	"fmt"
)

func main() {
	// continue打印1--100以内的奇数(要求使用for循环 + continue)
	for i := 1; i <= 100; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("i = ", i)
	}

	// 从键盘读入个数不确定的整数 并判断读入的整数和负数的个数
	// 输入为0时结束程序(使用for循环 continue break)完成
	var positiveCount int
	var negativeCount int
	var num int
	for {
		fmt.Println("请输入一个整数:")
		fmt.Scanln(&num)
		if num == 0 {
			break
		}
		if num > 0 {
			positiveCount++
			continue
		}
		negativeCount++
	}
	fmt.Println("正数的个数是:", positiveCount, " 负数的个数是:", negativeCount)
}

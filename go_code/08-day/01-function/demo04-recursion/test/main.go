package main

import "fmt"

func fibonacci(n int) int {
	if n == 1 || n == 2 {
		return 1
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func fx(n int) int {
	if n == 1 {
		return 3
	}
	return fx(n-1)*2 + 1
}

// 猴子吃桃子 有一堆桃子 猴子第一天吃了其中的一半 并再多吃了一个
// 以后每天猴子都吃其中的一般 然后再多吃一个
// 当到第十天的时候 想再吃时(还没有吃) 发现只有一个桃子了 问题 最初共多少个桃子
func monkeyPeach(n int) int {
	if n > 10 || n < 1 {
		fmt.Println("输入的天数不对")
		return 0
	}
	if n == 10 {
		return 1
	} else {
		return (monkeyPeach(n+1) + 1) * 2
	}
}

func main() {
	fmt.Println(fibonacci(1))
	fmt.Println(fibonacci(2))
	fmt.Println(fibonacci(3))
	fmt.Println(fibonacci(5))
	fmt.Println("-------------------")
	fmt.Println(fx(1))
	fmt.Println(fx(2))
	fmt.Println(fx(3))
	fmt.Println(fx(4))
	fmt.Println("-------------------")
	fmt.Println(monkeyPeach(1))
}

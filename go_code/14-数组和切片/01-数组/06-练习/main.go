package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 要求 随机生成5个数 并将其反转打印
	// 1 随机生成5个数 rand.Intn()函数
	// 2 当我们得到随机数后 就放到一个数组 int数组
	// 3 反转打印
	intArr := [6]int{}
	rand.Seed(time.Now().UnixNano())
	// 为了每次生成的随机数不一样 我们需要给一个seed值
	for i := 0; i < len(intArr); i++ {
		intArr[i] = rand.Intn(100) // 0<= n <100
	}
	fmt.Println(intArr)
	for i := len(intArr) - 1; i >= 0; i-- {
		fmt.Println(intArr[i])
	}
}

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 我们为了生成一个随机数 还需要给rand设置一个种子
	// time.Now().Unix() 返回一个从 1970:01:01的0时0分0秒 到现在的秒数
	// rand.Seed(time.Now().Unix())
	// 如何随机的生成0-100整数
	// n := rand.Intn(100) + 1 //[0,100)
	// fmt.Println(n)

	// 随机生成1-100的一个数 直到生成了99这个数 看看里一共用了几次
	i := 0
	for {
		rand.Seed(time.Now().UnixNano()) // 返回纳秒数
		n := rand.Intn(100) + 1
		fmt.Println(n)
		i++
		if n == 99 {
			break
		}
	}
	fmt.Println("一共用了",i,"次")
}

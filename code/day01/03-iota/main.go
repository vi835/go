package main

import (
	"fmt"
)

// iota
const (
	a1 = iota // 0
	a2        // 1
	a3        // 2
)

// 使用 _ 跳过某些值
const (
	b1 = iota // 0
	b2 = iota // 1
	_         // 2
	b3 = iota // 3
)

// 插队
const (
	c1 = iota // 0
	c2 = 100  // 100
	c3 = iota // 2
	c4 = iota // 3
)

// 多个常量声明在一行
const (
	d1, d2 = iota + 1, iota + 2 // 1,2
	d3, d4 = iota + 1, iota + 2 // 2,3
)

// 定义数量级
// 这里的<<表示左移操作，1<<10表示将1的二进制表示向左移10位，也就是由1变成了10000000000，
// 也就是十进制的1024。同理2<<2表示将2的二进制表示向左移2位，也就是由10变成了1000，也就是十进制的8
const (
	_  = iota
	KB = 1 << (10 * iota)
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)

func main() {
	fmt.Println("a1:", a1)
	fmt.Println("a2:", a2)
	fmt.Println("a3:", a3)

	fmt.Println("b1:", b1)
	fmt.Println("b2:", b2)
	fmt.Println("b3:", b3)

	fmt.Println("c1:", c1)
	fmt.Println("c2:", c2)
	fmt.Println("c3:", c3)
	fmt.Println("c4:", c4)

	fmt.Println("d1:", d1)
	fmt.Println("d2:", d2)
	fmt.Println("d3:", d3)
	fmt.Println("d4:", d4)
}

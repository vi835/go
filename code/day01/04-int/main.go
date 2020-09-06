package main

import (
	"fmt"
)

// 整形
func main() {
	var i1 = 101
	fmt.Printf("%d\n", i1)
	// 转成二进制输出101 1100101
	fmt.Printf("%b\n", i1)
	// 十进制数转成八进制
	fmt.Printf("%o\n",i1)
	// 十进制数转成十六进制
	fmt.Printf("%x\n",i1)

	// 八进制
	i2 := 077
	// 八进制数077 表示成十进制数位63
	fmt.Printf("%d\n", i2)

	// 十六进制
	i3 := 0xee123
	fmt.Printf("%d\n", i3)

	// 查看变量的类型
	fmt.Printf("%T\n",i1)

	// 声明int8类型的变量
	i4:=int8(9)
	fmt.Printf("%d\n", i4)
}

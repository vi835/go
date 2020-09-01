package main

import (
	"fmt"
)

func main() {
	// 位运算的演示
	fmt.Println(2 & 3) // 2
	fmt.Println(2 | 3) // 3
	fmt.Println(2 ^ 3) // 1

	// 2  的补码 0000 0010
	// -2 的原码 1000 0010 
	//     反码  1111 1101
	//     补码  1111 1110
	fmt.Println(2 & -2) // 2
	fmt.Println(2 | -2) // -2
	fmt.Println(2 ^ -2) // -4 1111 1100 || 1111 1011 || 1000 0100

	a := 1 >> 2 // 0000 0001 => 0000 0000 => 0
	c := 1 << 2 // 0000 0001 => 0000 0100 => 4
	fmt.Println("a = ",a,"c = ",c)
}

package main

import (
	"fmt"
)

// 使用switch把小写类型的char型转换为大写 至转换 a b c d e 其他的输出 other
func main()  {
	var char byte
	fmt.Println("请输入要转换的字符")
	fmt.Scanf("%c",&char)
	fmt.Println(char)
	switch char {
	case 'a','b','c','d','e':
		char -=32
		fmt.Printf("转换后的字符是:%c\n",char)
	default:
		fmt.Println("other")
	}
}
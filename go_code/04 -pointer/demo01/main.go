package main

import(
	"fmt"
)

func main()  {
	// 演示一把 * 和 & 的使用
	a:=100
	fmt.Println("a的地址是=====>",&a)

	var ptr *int = &a
	fmt.Println("ptr指向的值是======>",*ptr)
}
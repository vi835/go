package main

import (
	"fmt"
)

func main() {
	fmt.Println("请输入字符:")
	var key byte
	fmt.Scanf("%c", &key)
	switch key {
	case 'a':
		fmt.Println("星期一")
	default:
		fmt.Println("请输入正确的字符")
	}

}

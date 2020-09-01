package main

import (
	"fmt"
)

func main() {
	fmt.Println("请输入您的年龄:")
	var age int
	fmt.Scanln(&age)
	if age > 18 {
		fmt.Println("你年龄大于18 要对自己的行为负责")
	}else{
		fmt.Println("你的年龄不大 这次放过你了")
	}

	// golang支持在if中 直接定义一个变量 如下
	if sal := 20; sal > 18 {
		fmt.Println("test")
	}
}

package main

import (
	"fmt"
	"go_code/08-day/01-function/demo05-init/utils"
)

// 为了看到全局变量是先被初始化的 我们这里先写函数
var age = test()

// intit函数 通常可以在init函数中完成初始化操作
func test() int {
	fmt.Println("test()..") // 1
	return 90
}

func init() {
	fmt.Println("init()...") // 2
}

func main() {
	fmt.Println("main()...age=", age) // 3
	fmt.Println("name=",utils.Name)
	fmt.Println("age=",utils.Age)
}

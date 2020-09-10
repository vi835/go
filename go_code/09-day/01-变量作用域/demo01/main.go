package main

import "fmt"

// 函数外部声明/定义的变量叫全局变量 作用域整个包都有效
// 如果其首字母大写 则作用域再整个程序有效
var age int = 50
var Name string = "jack"

func test() {
	// age和 Name的作用域就只有再test函数内部
	age := 10
	Name := "tom"
	fmt.Println("age=", age)
	fmt.Println("Name=", Name)
}
func main() {
	fmt.Println("age=", age)
	fmt.Println("Name=", Name)

	// 如果变量是在一个代码块 比如for/if 那么这个变量的作用域就在该代码块
	for i := 0; i <= 10; i++ {
		fmt.Println("i=", i)
	}
}

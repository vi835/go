package main

// 导入语句
import (
	"fmt"
)

// 函数外只能放置标识符(变量/常量/函数/类型)的声明
// fmt.Println("hello") 非法

// 批量声明
var (
	name string
	age  int
	isOk bool
)

// 程序的入口函数
func main() {
	fmt.Println("Hello world")
	name = "理想"
	age = 16
	isOk = true
	// 在终端中输入要打印的内容
	fmt.Println(isOk)
	// %s 占位符 使用name这个变量的值去替换占位符
	fmt.Printf("name:%s\n", name)
	// 打印完指定的内容之后会在后面加上一个换行符
	fmt.Println(age)

	// 声明变量同时赋值
	var s1 string = "abc"
	fmt.Println(s1)
	// 类型推导(根据值判断变量变量是什么类型)
	var s2 = "hello"
	fmt.Println(s2)
	// 简短变量声明 只能在函数中使用
	s3 := "哈哈哈"
	fmt.Println(s3)
}

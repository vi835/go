package main

import (
	"fmt"
	"strings"
)

// 编写一个函数 可以接收一个文件后缀名 并返回一个闭包
// 调用闭包 可以传入一个文件名 如果该文件名没有指定的后缀 则返回文件名.后缀名
// 如果已经有指定后缀 则返回原文件名
func makeSuffix(suffix string) func(str string) string {
	return func(str string) string {
		if strings.HasSuffix(str, suffix) {
			return str
		}
		return str + suffix
	}
}
func main() {
	// 测试一下makeSuffix
	func1 := makeSuffix(".jpg")
	fmt.Println("文件名是=====>",func1("aaaa.jpg"))
	fmt.Println("文件名是=====>",func1("bbbb"))
}

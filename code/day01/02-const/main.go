package main

import (
	"fmt"
)

// 常量的使用
func main() {
	// 定义了常量之后不能修改
	// 在程序运行期间不能改变
	const pi = 3.1415926

	// 批量声明常量
	const (
		statusOk = 200
		notFound = 404
		// 批量声明常量时 如果某一行没有
		// 赋值 默认就和上一行一致
		err
	)
	fmt.Println("pi:", pi)
	fmt.Println("statusOk:", statusOk)
	fmt.Println("notFound:", notFound)
	fmt.Println("err:", err)
}

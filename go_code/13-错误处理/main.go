package main

import (
	"errors"
	"fmt"
)

func test() {
	// 使用 defer+recover 来捕获和处理异常
	defer func() {
		err := recover() // recover() 内置函数 可以捕获异常
		if err != nil {  // nil是err的0值 如果不等于0值 说明捕获到异常
			fmt.Println("err=", err)
			// 这里就可以进行错误处理
		}
	}()
	num1 := 10
	num2 := 0
	res := num1 / num2
	fmt.Println("res=", res)
}

// 函数去读取以配置文件init.conf的信息
// 如果文件名传入不正确 我们就返回一个自定义的错误
func readConf(name string) (err error) {
	if name == "config.ini" {
		// 读取...
		return nil
	} else {
		// 返回一个自定义错误
		return errors.New("读取文件错误...")
	}
}

func test2() {
	err :=readConf("confhig.ini")
	if err !=nil{
		// 如果读取文件错误 就输出这个错误 并终止程序
		panic(err)
	}
	fmt.Println("test2()继续执行")
}
func main() {
	test()
	fmt.Println("main()中函数")

	// 测试自定义错误的使用
	test2()
}

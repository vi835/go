package main

import (
	"fmt"
)

func main() {
	// 说明 如果运算的数都是整数 那么除后 去掉小数部分 保留整数部分
	fmt.Println(10 / 4)

	var n1 float32 = 10 / 4
	fmt.Println(n1)

	// 如果我们希望保留小数部分 则需要有浮点数参与运算
	var n2 float32 = 10.0 / 4
	fmt.Println(n2)

	// 演示 % 的使用
	// 看一个公式 a%b = a - a/b*b
	fmt.Println("10%3=====>", 10%3)
	fmt.Println("-10%3=====>", -10%3) // -10 -(-10)/3*3
	fmt.Println("10%-3=====>", 10%-3)
	fmt.Println("-10%-3=====>", -10%-3)

	// ++ 和 --的使用
	var i int = 10
	i++ // <=> i=i+1
	fmt.Println("i=", i)
	i--
	fmt.Println("i=", i)

	fmt.Printf("%v个星期零%v天\n", 97/7, 97%7)

	// 有两个变量 a和b 要求将其进行交换 但是不允许使用中间变量 最终打印结果
	var a = 10
	var b = 20

	a = a + b
	b = a - b
	a = a - b

	fmt.Println("a===>",a)
	fmt.Println("b===>",b)
}

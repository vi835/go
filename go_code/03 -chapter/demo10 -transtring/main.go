package main

import (
	"fmt"
	"strconv"
)

// 演示golang中基本数据类型转成string使用
func main() {
	var num1 int = 90
	var num2 float64 = 23.456
	var b bool = true
	var mychar byte = 'h'
	var str string // 空的string

	// 使用第一种方式来转换 fmt.Sprintf方法
	str = fmt.Sprintf("%d", num1)
	fmt.Printf("str type %T str = %v\n", str, str)

	str = fmt.Sprintf("%f", num2)
	fmt.Printf("str type %T str = %v\n", str, str)

	str = fmt.Sprintf("%t", b)
	fmt.Printf("str type %T str = %v\n", str, str)

	str = fmt.Sprintf("%c", mychar)
	fmt.Printf("str type %T str = %v\n", str, str)

	// 第二种方式 strconv函数
	var num3 int = 99
	var num4 float64 = 23.456
	var b2 bool = true

	// 10代表10进制
	str = strconv.FormatInt(int64(num3), 10)
	fmt.Printf("str type %T str = %v\n", str, str)

	// 说明 'f'代表格式 10 表示小数点位保留10位 64表示这个小数是float64
	str = strconv.FormatFloat(num4, 'f', 10, 64)
	fmt.Printf("str type %T str = %v\n", str, str)

	str = strconv.FormatBool(b2)
	fmt.Printf("str type %T str = %v\n", str, str)

	// strconv包中有一个函数 Itoa
	var num5 int = 4567
	str = strconv.Itoa(num5)
	fmt.Printf("str type %T str = %v\n", str, str)
}

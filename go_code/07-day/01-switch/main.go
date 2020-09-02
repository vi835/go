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

	// switch后也可以不带表达式 类似 if-else分支来使用
	var age int = 10
	switch {
	case age == 10:
		fmt.Println("age == 10")
	case age == 20:
		fmt.Println("age == 20")
	default:
		fmt.Println("没有匹配到")
	}

	// switch 穿透 fallthrough
	var num int = 10
	switch num {
	case 10:
		fmt.Println("ok1")
		fallthrough // 默认只能穿透一层
	case 20:
		fmt.Println("ok2")
	case 30:
		fmt.Println("ok3")
	default:
		fmt.Println("没有匹配到")
	}

	// case 后面的表达式如果是常量值 则要求不能重复
	var n1 int32 = 5
	var n2 int32 = 20
	switch n1 {
	case n2, 10, 5: // case 后面可以有多个表达式
		fmt.Println("ok1")
	// case 5: // 错误 因为前面我们有常量5 因此重复 就会报错
	// 	fmt.Println("ok2")
	default:
		fmt.Println("没有匹配到")
	}

}

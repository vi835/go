package main

// fmt包中提供格式化 输出 输入的函数
import "fmt"

func main() {
	// 演示转义字符的使用
	// 制表符的使用 \t
	fmt.Println("di\tdi")

	// 换行符的使用 \n
	fmt.Println("hello\nworld")

    // \的转义
	fmt.Println("D:\\学习笔记\\GO\\笔记")

	// "的转义
	fmt.Println("tom 说\"I love you\"")

	// /r 回车  从当前行的最前面开始输出 覆盖掉以前的内容
	fmt.Println("天龙八部 雪山飞狐\r张飞")

	// 练习
	fmt.Println("姓名\t年龄\t籍贯\t住址\njohn\t12\t河北\t北京\t")
}

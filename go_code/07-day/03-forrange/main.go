package main

import "fmt"

func main() {
	// 字符串遍历1  传统方式 使用切片 解决中文乱码问题
	var str string = "hello,你好"
	str2 :=[]rune(str)
	for i := 0; i < len(str2); i++ {
		fmt.Printf("%c\n", str2[i])
	}

	// 字符串遍历方式2 for-range 是按照字符的方式进行遍历的
	str = "abcok,你好"
	for index,val:=range str{
		fmt.Printf("index = %d ,val = %c \n",index,val)
	}
}

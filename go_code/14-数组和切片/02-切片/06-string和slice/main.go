package main

import "fmt"

func main() {
	// string底层是一个byte数组 因此string也可以进行切片处理
	str := "hello@atguigu"
	// 使用切片获取到 atguigu
	slice := str[6:]
	fmt.Println("slice=", slice)

	// string是不可变的 也就是说不能通过str[0]='z' f方式来修改字符串
	// str[0]='z'
	// 如果需要修改字符串 可以先将string->byte[] 或者 []rune->修改 重写string
	// "hello@atguigu"->"zello@atguigu"
	arr1 := []byte(str)
	arr1[0] = 'z'
	str = string(arr1)
	fmt.Println("str=", str)

	// 细节 我们转成[]byte后 可以处理英文和数字 但是不能处理中文
	// 原因 []byte 按字节来处理 而一个汉字 是三个字节 因此会出现乱码
	// 解决方法是将 string转成 []rune即可 因为rune是按字符处理的 可以兼容汉字
	arr2 := []rune(str)
	arr2[0] = '北'
	str = string(arr2)
	fmt.Println("str=", str)
}

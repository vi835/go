package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// 统计字符串的长度 按字节(len())
	str := "hello"
	// golang的编码统一为utf-8(ascii的字符(字母和数字)占一个字节 汉字占用三个字节)
	str1 := "hello你好1"
	fmt.Println("str的长度是:", len(str))
	fmt.Println("str1的长度是:", len(str1))

	// 字符串遍历 同时出差有中文的问题 r:[] = []rune(str2)
	r := []rune(str1)
	for i := 0; i < len(r); i++ {
		fmt.Printf("字符有%c\n", r[i])
	}

	// 字符串转整数 n,err := strconv.Atoi("12 )
	n, err := strconv.Atoi("123")
	if err != nil {
		fmt.Println("转换错误", err)
	} else {
		fmt.Println("转换的结果是:", n)
	}

	// 整数转字符串  strconv.Itoa(12345)
	str = strconv.Itoa(12345)
	fmt.Printf("str=%v,str=%T\n", str, str)

	// 字符串 转 []byte  var bytes =[]byte("hello world")
	var bytes = []byte("hello world")
	// 输出对应字符的编码值
	fmt.Printf("bytes=%v\n", bytes)

	//  []byte 转 字符串 str = string([]byte{97, 98, 99})
	str = string([]byte{97, 98, 99})
	fmt.Printf("str=%v\n",str)

	// 10进制转 2 8 16进制 str = strconv.FormatInt(123,2) 后面一个参数是进制
	str = strconv.FormatInt(123,2)
	fmt.Println("123对应的二进制是:%v\n",str)
	str = strconv.FormatInt(123,16)
	fmt.Println("123对应十六进制是:%v\n",str)

	// 查找子串是否在指定的字符串中 strings.Contains("seafood","sea")
	b:=strings.Contains("seafood","sea")
	fmt.Printf("b=%v\n",b)

	// 统计一个字符串有几个指定的字串 strings.Count("ceheese","e")
	num := strings.Count("ceheese","e")
	fmt.Printf("num=%v\n",num)

	// 字符串比较  == 区分大小写
	// strings.EqualFold("abc","aBC") 不区分大小写
	b = "abc"=="Abc"
	fmt.Printf("b=%v\n",b)
	b = strings.EqualFold("abc","aBC")
	fmt.Printf("b=%v\n",b)

	// 返回子串在字符串中第一次出现的index值 如果没有返回-1
	index:=strings.Index("NLT_abc","abc")
	fmt.Printf("index=%v\n",index)
}

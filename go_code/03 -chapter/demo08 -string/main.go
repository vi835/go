package main

import (
	"fmt"
)

// 演示Golang中string类型的使用
func main() {
	// string的使用
	var address string = "北京长城 110 hello world"
	fmt.Println(address)

	// 字符串一旦赋值了 字符串就不能修改了 在Go中字符串是不可变的
	var str = "hello"
	// str[0] = 'a' // 这里就不能去修改str的内容 即go中的字符串是不可变的
	fmt.Println(str)

	// 字符串的两种表示形式
	// 双引号 会识别转义字符
	// 反引号 以字符串的原生形式输出 包括换行和特殊字符 可以实现防止攻击 输出源代码等效果
	str2 := "abc\nhello"
	fmt.Println(str2)

	// 使用的反引号 ``
	str3 := `// 演示Golang中string类型的使用
	func main(){
		// string的使用
		var address string = "北京长城 110 hello world"
		fmt.Println(address)
	}`
	fmt.Println(str3)

	// 字符串拼接方式
	var str4 = "hello" + "world"
	str4 += " haha!"
	fmt.Println(str4)

	// 当一个拼接的操作很长时 怎么办 可以分行写
	// 但是 注意 需要将 + 保留在上一行
	var str5 = "hello" + " world" +
		"hello" + " world" +
		"hello" + " world"
	fmt.Println(str5)

	var a int          // 0
	var b float32      // 0
	var c float64      // 0
	var isMarried bool // false
	var name string    // ""
	// 这里的v 表示按照变量的值输出
	fmt.Printf("a = %d,b = %v,c = %v,isMarried = %v,name = %v", a, b, c, isMarried, name)
}

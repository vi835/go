package main

import (
	"fmt"
	"unsafe"
)

// 演示Golang中bool类型使用
func main() {
	var b = false
	fmt.Println("b =", b)
	// 注意事项
	// bool类型占用存储空间是一个字节
	fmt.Println("bool类型占用的空间大小为 ", unsafe.Sizeof(b))
}

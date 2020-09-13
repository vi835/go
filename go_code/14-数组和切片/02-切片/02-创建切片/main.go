package main

import "fmt"

func main() {
	// 演示切片的使用 make方法
	var slice []float64=make([]float64,5,10)
	// 对于切片 必须make后使用
	fmt.Println(slice)
	fmt.Println("slice的len",len(slice))
	fmt.Println("slice的cap",cap(slice))

	// 第三种方式 定义一个切片 直接就指定具体数组 使用原理类型make的方式
	var strSlice []string =[]string {"tom","jack","mary"}
	fmt.Println(strSlice)
	fmt.Println("strSlicee的len",len(strSlice))
	fmt.Println("strSlice的cap",cap(strSlice))
}

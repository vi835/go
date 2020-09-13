package main

import "fmt"

// Go的数组属于值类型 在默认情况下是值传递 因此会进行值拷贝 数组间不会相互影响
func test01(arr [3]int) {
	arr[0] = 111
	fmt.Println("test01里的arr=", arr)
}

// 如果想在其他函数中 去修改原来的数组 可以使用引用传递(指针)
func test02(arr *[3]int) {
	(*arr)[0] = 88
	fmt.Println("test02里的arr=", *arr)
}
func main() {
	// 数组是多个相同类型数据的组合 一个数组一旦声明/定义了 其长度是固定的 不能动态改变
	var arr01 [3]int
	arr01[0] = 1
	arr01[1] = 2
	arr01[2] = 3
	// arr01[2] = 2.1 // 这里会报错 类型不对
	// arr01[3] = 4 //长度是笃定的 不能动态改变 否则会报错

	arr := [3]int{0, 1, 2}
	test01(arr)
	fmt.Println(arr)

	test02(&arr)
	fmt.Println(arr)
}

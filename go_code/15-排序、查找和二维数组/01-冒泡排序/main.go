package main

import "fmt"

func BubbleSort(arr *[5]int) {
	fmt.Println("排序前arr=", *arr)
	temp := 0
	for j := 0; j < len(*arr)-1; j++ {
		for i := 0; i < len(*arr)-1-j; i++ {
			if (*arr)[i] > (*arr)[i+1] {
				temp = (*arr)[i]
				(*arr)[i] = (*arr)[i+1]
				(*arr)[i+1] = temp
			}
		}
	}
}

func test(arr *[5]int) {
	temp := 0
	fmt.Println(*arr)
	for i := 0; i < len(*arr)-1; i++ {
		for j := 0; j < len(*arr)-i-1; j++ {
			if (*arr)[j] < (*arr)[j+1] {
				temp = (*arr)[j]
				(*arr)[j] = (*arr)[j+1]
				(*arr)[j+1] = temp
			}
		}
	}
}
func main() {
	// 定义数组
	arr := [...]int{24, 69, 80, 57, 13}
	// 将数组传递给一个函数 完成排序
	//BubbleSort(&arr)
	test(&arr)
	fmt.Println(arr)
}

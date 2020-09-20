package main

import "fmt"

func main() {
	// 定义/声明二维数组
	var arr [4][6]int

	// 赋初值
	arr[1][2] = 1
	arr[2][1] = 2
	arr[2][3] = 3

	// 遍历二维数组 按照要求输出图形
	for i := 0; i < 4; i++ {
		for j := 0; j < 6; j++ {
			fmt.Print(arr[i][j], " ")
		}
		fmt.Println()
	}
	fmt.Println(arr)
}

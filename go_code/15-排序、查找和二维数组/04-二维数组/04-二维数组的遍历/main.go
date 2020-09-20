package main

import "fmt"

func main() {
	// 演示二维数组的遍历
	var arr3 = [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println("arr3=", arr3)

	// 双重for循环遍历
	for i := 0; i < len(arr3); i++ {
		for j := 0; j < len(arr3[i]); j++ {
			fmt.Print(arr3[i][j], " ")
		}
		fmt.Println()
	}

	fmt.Println()
	// for-range遍历
	for i, v := range arr3 {
		for m, n := range v {
			fmt.Printf("arr3[%v][%v]=%v \t", i, m, n)
		}
		fmt.Println()
	}
}

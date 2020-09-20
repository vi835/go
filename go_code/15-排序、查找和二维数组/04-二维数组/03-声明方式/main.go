package main

import "fmt"

func main() {
	var arr1 [2][3]int = [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println("arr1=", arr1)

	var arr2 [2][3]int = [...][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println("arr2=", arr2)

	var arr3 = [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println("arr3=", arr3)

	var arr4 = [...][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println("arr4=", arr4)
}

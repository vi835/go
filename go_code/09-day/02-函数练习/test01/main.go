package main

import "fmt"

func printFunc(num int) {
	for i := 1; i <= num; i++ {
		for k := num - i; k >= 0; k-- {
			fmt.Print(" ")
		}
		for j := 1; j < 2*i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func main() {
	fmt.Println("请输入一个整数:")
	var num int
	fmt.Scanln(&num)
	fmt.Println(num)
	printFunc(num)
}

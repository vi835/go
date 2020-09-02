package main

import "fmt"

func main() {
	// 打印1-100之间所有是9的倍数的整数的个数以及总和
	var start int = 9
	count := 0
	sum := 0
	for {
		if start > 100 {
			break
		}
		sum += start
		start += 9
		count++
	}
	fmt.Println("个数是",count,"个，总和是",sum)
}

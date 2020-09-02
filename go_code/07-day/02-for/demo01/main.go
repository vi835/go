package main

import (
	"fmt"
)

func main() {
	for i := 1; i < 11; i++ {
		fmt.Println("hello", i)
	}

	// for循环的第二种写法
	j := 1
	for j <= 10 {
		fmt.Println("hello", j)
		j++
	}

	// for循环的第三种写法
	k := 1
	for {
		if k <= 10 {
			fmt.Println("hello")
		} else {
			break
		}
		k++
	}
}

package main

import "fmt"

func main() {
	// 演示for-range遍历数组
	heroes := [...]string{"宋江", "卢俊义", "吴用"}
	fmt.Println(heroes)

	for i, v := range heroes {
		fmt.Printf("index=%v value=%v\n", i, v)
	}

	for _, v := range heroes {
		fmt.Printf("元素的值是%v\n", v)
	}
}

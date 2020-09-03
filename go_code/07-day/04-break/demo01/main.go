package main

import (
	"fmt"
)

func main() {

	// 这里演示一下指定标签的形式来使用break
label2:
	for i := 0; i < 4; i++ {
		// label1:
		for j := 0; i < 10; j++ {
			if j == 2 {
				// break // break 默认会默认跳出最近的for循环
				break label2
			}
			fmt.Println("j = ", j)
		}
	}
}

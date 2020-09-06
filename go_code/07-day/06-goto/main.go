package main

import (
	"fmt"
)

func main(){
	// 演示goto的使用
	fmt.Println("ok1")
	goto label1
	fmt.Println("ok2")
	fmt.Println("ok3")
	label1:
	fmt.Println("ok4")
	fmt.Println("ok5")
}


package main

import(
	"fmt"
)

func main(){
	var num int = 9
	fmt.Printf("num 地址=====> %v\n",&num)

	var ptr *int
	ptr = &num
	*ptr = 10 //这里修改值 会引起num的值变化
	fmt.Printf("num 值=====> %v\n",num)
}
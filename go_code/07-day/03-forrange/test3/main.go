package main

import(
	"fmt"
)

func main(){
	// 使用 for循环 打印一个金字塔
	var floor int 
	fmt.Println("请输入要打印的金字塔层数:")
	fmt.Scanln(&floor)

	for i := 1; i <= floor; i++ {
		for j := 1; j <= floor+i; j++ {
			if j <= floor-i+1 {
				fmt.Print(" ")
			} else {
				fmt.Print("*")
			}
		}
		fmt.Println()
	}

	fmt.Println()

	for i:=1;i<=floor;i++{
		for j:=floor-i;j>=1;j--{
			fmt.Print(" ")
		}
		for k:=1;k<=i*2-1;k++{
			fmt.Print("*")
		}
		fmt.Println()
	}

	// 打印九九乘法表
	for i:=1;i<=9;i++{
		for k:=1;k<=i;k++{
			fmt.Printf("%d * %d = %d ",k,i,i*k)
		}
		fmt.Println()
	}


}
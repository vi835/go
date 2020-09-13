package main

import "fmt"

func main() {
	var intArr [3]int
	// 当我们定义完数组后 其实数据的各个元素有默认值0
	fmt.Println(intArr)
	// 地址是16进制
	// 打印数组的地址就是第一个元素的地址 int数组的下一个元素的地址=上一个元素的地址+int占用字节数(8)
	fmt.Printf("intArr的地址=%p intArr[0]的地址=%p\n",&intArr,&intArr[0])
	fmt.Println(&intArr[0])
	fmt.Println(&intArr[1])
	fmt.Println(&intArr[2])

	// 从终端循环输入5个成绩 保存到float64数组 并输出
	var scores [5]float64
	for i:=0;i<len(scores);i++{
		fmt.Println("请输入成绩:")
		fmt.Scanln(&scores[i])
	}
	fmt.Println(scores)

}

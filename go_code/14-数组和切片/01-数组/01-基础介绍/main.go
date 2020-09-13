package main

import "fmt"

func main() {
	// 使用数组的方式来解决问题

	// 1 定义一个数组
	var hens [6]float64
	// 2 给数组中的每一个元素赋值 数组元素的下标是从0开始的
	hens[0] = 3.0
	hens[1] = 5.0
	hens[2] = 2.0
	hens[3] = 3.4
	hens[4] = 3.0
	hens[5] = 50.0
	// 3 遍历数组求出总体重
	totalWeight := 0.0
	for i := 0; i < len(hens); i++ {
		totalWeight += hens[i]
	}
	// 求出平均体重
	// totalWeight是float64类型 如果需要用变量进行算数 需要进行强转
	avgWeight := fmt.Sprintf("%.2f", totalWeight/6)
	avgWeight = fmt.Sprintf("%.2f", totalWeight/float64(len(hens)))
	fmt.Printf("totalWeight=%v avgWeight=%v\n", totalWeight, avgWeight)

	// 四种初始化数组的方式
	var numArr1 [3]int = [3]int{1, 2, 3}
	fmt.Println("numArr1=", numArr1)

	var numArr2 = [3]int{2, 3, 4}
	fmt.Println("numArr2=", numArr2)

	// 这里的[...]是固定写法
	var numArr3 = [...]int{3, 4, 5, 6}
	fmt.Println("numArr3=", numArr3)

	// 初始化数组赋值的同时 可以指定复制元素顺序
	var numArr4 = [...]int{1: 800, 0: 900, 2: 999}
	fmt.Println("numArr4=", numArr4)

	// 类型推导
	arr5 := [...]string{"tom", "jerry", "hong"}
	fmt.Println("arr5=", arr5)
}

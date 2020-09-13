package main

import "fmt"

func main() {
	characters := [26]byte{}
	for i := 0; i < len(characters); i++ {
		characters[i] = 'A' + byte(i)
	}
	for _, v := range characters {
		fmt.Printf("%c\n", v)
	}
	fmt.Printf("characters=%v\n", characters)

	// 声明一个数组
	intArr := [...]int{1, -1, 9, 9, 90, 12}
	max := [...]int{0, intArr[0]}
	sum := 0
	for i, v := range intArr {
		sum += v
		if max[1] < v {
			max[0] = i
			max[1] = v
		}
	}
	fmt.Printf("数组最大值是%v 下标是%v\n", max[1], max[0])
	// 如何让平均值保留到小数
	fmt.Printf("sum=%v avg=%v",sum,float64(sum)/float64(len(intArr)))
}

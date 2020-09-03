package main

import (
	"fmt"
)

func main() {
	// 统计三个班成绩情况 每个班有5名同学
	// 求出各个班的平均分和所有班级的平均分(学生的成绩从键盘输入)
	var average float32
	var score float32
	var sum float32
	for i := 0; i < 3; i++ {
		sum = 0
		for j := 0; j < 5; j++ {
			fmt.Println("请输入学生的成绩:")
			fmt.Scanln(&score)
			sum += score
		}
		average += sum / 5
		fmt.Println("第", i+1, "个班的平均成绩是:", sum/5)
	}
	fmt.Println("三个班的平均成绩是:", average/3)
}

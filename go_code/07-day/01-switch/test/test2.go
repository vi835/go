package main

import(
	"fmt"
)

// 对学生成绩大于60分的 输出合格 低于60分的 输出不合格 学生的成绩不大于100
func main1() {
	fmt.Println("请输入学生的成绩:")
	var score float32
	fmt.Scanln(&score)
	switch int(score/60) {
	case 0:
		fmt.Println("不合格")
	default:
		fmt.Println("合格")
	}
}
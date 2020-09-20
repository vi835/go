package main

import "fmt"

func main() {
	// 有一个数列 白眉鹰王 金毛狮王 紫衫龙王 青翼蝠王
	// 猜数游戏 从键盘中任意输入一个名称 判断数列中是否包含此名称

	names := [...]string{"白眉鹰王", "金毛狮王", "紫衫龙王", "青翼蝠王"}
	heroName := ""
	fmt.Println("请输入要查找的人...")
	fmt.Scanln(&heroName)
	// 顺序查找 第一种方式
	for i := 0; i < len(names); i++ {
		if heroName == names[i] {
			fmt.Printf("找到了%v 下标是%v\n", heroName, i)
			break
		} else if i == (len(names) - 1) {
			fmt.Printf("没有找到%v\n", names)
		}
	}

	// 顺序查找 第二种方式
	index := -1
	for i := 0; i < len(names); i++ {
		if heroName == names[i] {
			index = 1
			break
		}
	}
	if index == -1{
		fmt.Printf("没有找到%v\n", names)
	}else{
		fmt.Printf("找到了%v 下标是%v\n", heroName, index)
	}
}

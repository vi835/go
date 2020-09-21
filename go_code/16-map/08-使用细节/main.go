package main

import "fmt"

func modify(map1 map[int]int) {
	map1[10] = 900
}

// 定义学生结构体
type Stu struct {
	Name    string
	Age     int
	Address string
}

func main() {
	// map是引用类型 遵守引用类型传递的机制 在一个函数接收map 修改后 会直接修改原来的map
	map1 := make(map[int]int)
	map1[1] = 90
	map1[2] = 88
	map1[10] = 10
	map1[20] = 20
	fmt.Println(map1)

	modify(map1)
	// 证明map是引用类型
	fmt.Println(map1)

	students := make(map[string]Stu, 10)
	// 创建两个学生
	stu1 := Stu{
		"Tom",
		18,
		"北京",
	}

	stu2 := Stu{"Jack",
		19,
		"上海",
	}
	students["n01"] = stu1
	students["n02"] = stu2
	fmt.Println(students)

	// 遍历各个学生的信息
	for k, v := range students {
		fmt.Printf("学生的编号是%v \n", k)
		fmt.Printf("学生的名字是%v \n", v.Name)
		fmt.Printf("学生的年龄是%v \n", v.Age)
		fmt.Printf("学生的地址是%v \n", v.Address)
	}
}

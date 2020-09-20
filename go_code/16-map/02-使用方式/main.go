package main

import "fmt"

func main() {
	// 1 第一种使用方式 先声明 再分配内存
	var a map[string]string
	a = make(map[string]string, 10)
	a["no1"] = "宋江"
	a["no2"] = "卢俊义"
	a["no3"] = "吴用"
	fmt.Println(a)

	// 第二种方式 直接分配内存
	cities := make(map[string]string, 10)
	cities["no1"] = "北京"
	cities["no2"] = "天津"
	cities["no3"] = "上海"
	fmt.Println(cities)

	// 第三种方式 声明直接赋值
	heroes := map[string]string{
		"hero1": "宋江",
		"hero2": "卢俊义",
	}
	fmt.Println(heroes)

	// 我们要存放3个学生信息 每个学生有name和sex信息
	studentMap := make(map[string]map[string]string)
	studentMap["stu01"] = make(map[string]string, 2)
	studentMap["stu01"]["name"] = "Tom"
	studentMap["stu01"]["sex"] = "男"

	studentMap["stu02"] = make(map[string]string, 2)
	studentMap["stu02"]["name"] = "Jack"
	studentMap["stu02"]["sex"] = "男"
	fmt.Println(studentMap)
	fmt.Println(studentMap["stu02"])
	fmt.Println(studentMap["stu01"]["name"])
}

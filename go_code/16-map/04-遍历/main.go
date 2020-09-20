package main

import "fmt"

func main() {
	// 说明 map的遍历使用for-range的结构遍历
	cities := make(map[string]string, 10)
	cities["no1"] = "北京"
	cities["no2"] = "天津"
	cities["no3"] = "上海"
	fmt.Println(cities)

	for k, v := range cities {
		fmt.Printf("key=%v value=%v \n", k, v)
	}

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

	for k1, v1 := range studentMap {
		fmt.Println("k1=", k1)
		for k2, v2 := range v1 {
			fmt.Printf("\t k2=%v v2=%v\n", k2, v2)
		}
	}
}

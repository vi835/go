package main

import "fmt"

func main() {
	// 说明 map的遍历使用for-range的结构遍历
	cities := make(map[string]string, 10)
	cities["no1"] = "北京"
	cities["no2"] = "天津"
	cities["no3"] = "上海"
	fmt.Println(cities)

	fmt.Println("cities 有",len(cities),"对")
}


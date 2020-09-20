package main

import "fmt"

func main() {
	// 第二种方式 直接分配内存
	cities := make(map[string]string, 10)
	cities["no1"] = "北京"
	cities["no2"] = "天津"
	cities["no3"] = "上海"
	fmt.Println(cities)

	// 因为key no3已经存在了 所以就是修改
	cities["no3"] = "上海~"
	fmt.Println(cities)

	// 演示删除
	delete(cities, "no3")
	// 当delete指定的key不存在时 删除不会操作 也不会报错
	delete(cities, "no4")
	fmt.Println(cities)

	// 演示map的查找
	val,ok:=cities["no1"]
	if ok {
		fmt.Println("有no1 key 它的值是=",val)
	}else {
		fmt.Println("没有no1key")
	}

	// 如果我们要删除map的所有key 没有一个专门的方法一次删除 可以遍历一下key 逐个删除
	// 或者map=make(...) make一个新的 让原来的称为垃圾 被gc回收
	cities = make(map[string]string, 10)
	fmt.Println(cities)

}

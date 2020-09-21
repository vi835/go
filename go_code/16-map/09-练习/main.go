package main

import "fmt"

// 使用 map[string]map[string]string 的map类型
// key 表示用户名 是唯一的 不可以重复
// 如果某个用户名存在 就将其密码修改"88888888" 如果不存在就增加这个用户信息
// 编写一个函数modifyUser(users map[string]map[string]string, name string)完成上述功能
func modifyUser(users map[string]map[string]string, name string) {
	// 判断users中是否有name
	// v, ok := users[name]
	if users[name] != nil {
		// 有这个用户
		users[name]["pws"] = "88888888"
	} else {
		// 没有这个用户
		users[name] = make(map[string]string, 2)
		users[name]["pwd"] = "88888888"
		users[name]["nickName"] = "昵称~" + name
	}
}
func main() {
	users := make(map[string]map[string]string, 10)
	modifyUser(users, "Tom")
	modifyUser(users, "mary")
	fmt.Println(users)
}

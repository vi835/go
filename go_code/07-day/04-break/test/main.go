package main

import (
	"fmt"
)

func main() {
	var sum int = 0
	for i := 1; i < 100; i++ {
		sum += i
		if sum > 20 {
			fmt.Println("和是:", sum)
			break
		}
	}
	var chances int = 3
	var username string
	var password string
	
	// 实现登录验证 有三次机会 登录成功 提示登录成功 
	// 否则提示还有几次机会
	for i:=1;i<=3;i++ {
		fmt.Println("请输入您的用户名密码:")
		fmt.Scanf("%s %s \n",&username,&password)
		if username == "张无忌" && password == "8888" {
			fmt.Println("登录成功")
			break
		}else{
			chances--
			fmt.Println("您剩下机会次数是:",chances)
		}
	}
}

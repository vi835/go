package main

import(
	"fmt"
	// 为了引用utils.go文件的变量或者函数 我们需要先引入该model包
	"go_code/03-chapter/demo12-variable/model"
)

func main()  {
	// 我们使用utils.go 的heroName 包名.标识符
	fmt.Println(model.HeroName)
}
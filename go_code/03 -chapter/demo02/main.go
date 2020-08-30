package main

import "fmt"

func main(){
	//  该案例演示golang如何一次性声明多个变量
    // var n1,n2,n3 int
	// fmt.Println("n1 =",n1,"n2 =",n2,"n3 =",n3)

	// 一次性声明多个变量的方式
	// var n1, name, n3 =100 ,"Tom" , 888
	// fmt.Println("n1 =",n1,"name =",name,"n3 =",n3)

	// 一次性声明多个变量的方式 同样可以使用类型推导
	n1, name, n3 :=100 ,"Tom" , 888
	fmt.Println("n1 =",n1,"name =",name,"n3 =",n3)
}
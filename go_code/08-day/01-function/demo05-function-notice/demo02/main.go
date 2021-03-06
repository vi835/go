package main

import (
	"fmt"
)

// 在Go中 函数也是一种数据类型
// 可以赋值给一个变量 则该变量就是一个函数类型的变量了 通过该变量可以对函数进行调用
func getSum(n1 int, n2 int) int {
	return n1 + n2
}

// 函数既然是一种数据类型 因此在Go中 函数可以作为形参 并且调用
func myFunc(funvar func(int, int) int, num1 int, num2 int) int {
	return funvar(num1, num2)
}

// 在加一个类型
// 这时 myFun 就是 func(int,int) int 类型
type myFunType func(int, int) int

func myFunc2(funvar myFunType, num1 int, num2 int) int {
	return funvar(num1, num2)
}

// 支持对函数返回值命名
func getSumAndSub(n1 int, n2 int) (sum int, sub int) {
	sum = n1 + n2
	sub = n1 - n2
	return
}
func main() {
	a := getSum
	fmt.Printf("a的数据类型%T,getSum的数据类型是%T\n", a, getSum)

	res := a(10, 40) // <=> res :=getSum(10,40)
	fmt.Println("res : ", res)
	fmt.Println("--------------")
	res2 := myFunc(getSum, 50, 60)
	fmt.Println("res2:", res2)
	fmt.Println("--------------")
	// 给int取了别名 在go中 myInt和int
	// 虽然都是int类型 但是go认为myInt和int是两种数据类型
	type myInt int
	var num1 myInt
	num1 = 40
	var num2 int = 20
	// num2 = int(num1)
	fmt.Println("num1 = ", num1)
	fmt.Printf("myInt的数据类型:%T,int的数据类型是%T\n", num1, num2)

	// 看案例
	res3 := myFunc2(getSum, 500, 600)
	fmt.Println("res3=", res3)

	m, n := getSumAndSub(10, 20)
	fmt.Println("m=", m, ",n=", n)
}

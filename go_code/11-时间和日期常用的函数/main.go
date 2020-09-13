package main

import (
	"fmt"
	"strconv"
	"time"
)

func test() {
	str := ""
	for i := 0; i < 10000; i++ {
		str += "hello" + strconv.Itoa(i)
	}
}

func main() {
	// 看看日期和时间相关函数和方法使用
	// 1 获取当前时间
	now := time.Now()
	fmt.Printf("now=%v now type=%T\n", now, now)

	// 2 通过now可以获取到年月日 时分秒
	// 通过int强转可以将英文月份转为数字
	fmt.Printf("年=%v\n", now.Year())
	fmt.Printf("月=%v\n", now.Month())
	fmt.Printf("月=%v\n", int(now.Month()))
	fmt.Printf("日=%v\n", now.Day())
	fmt.Printf("时=%v\n", now.Hour())
	fmt.Printf("分=%v\n", now.Minute())
	fmt.Printf("秒=%v\n", now.Second())

	// 3 格式化日期时间 Sprintf打印的同时返回一个字符串 使用格式化输出
	fmt.Printf("当前年月日 %d-%d-%d %d:%d:%d \n", now.Year(), now.Month(),
		now.Day(), now.Hour(), now.Minute(), now.Second())
	dateStr := fmt.Sprintf("当前年月日 %d-%d-%d %d:%d:%d \n", now.Year(), now.Month(),
		now.Day(), now.Hour(), now.Minute(), now.Second())
	fmt.Println("dateStr=", dateStr)

	// 格式化日期的第二种方式 格式化日期格式必须这么写
	// "2006-01-02 15:04:05" 标识符可以随意组合 日期必须是这几个
	fmt.Println(now.Format("2006-01-02 15:04:05"))
	fmt.Println(now.Format("2006-01-02"))
	fmt.Println(now.Format("15:04:05"))

	// 每隔一秒钟 打印一个数字 打印到10时就退出
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
		// 休眠
		// time.Sleep(time.Second)
	}

	// Unix和UnixNano的使用
	fmt.Printf("unix时间戳=%v unixNano时间戳%v\n", now.Unix(), now.UnixNano())

	// 先获取到当前的unix时间戳
	start := time.Now().UnixNano()
	test()
	end := time.Now().UnixNano()
	fmt.Printf("执行test函数花费的时间是%v纳秒\n", end-start)
}

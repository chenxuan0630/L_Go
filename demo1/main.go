package main

import "fmt"

// 全局变量可以采用一次性声明
var (
	username = "xuan"
	password = 159951
)

func main() {
	// 变量练习
	var test int
	fmt.Println(test) //输出0

	// 省略var
	sex := "男"
	fmt.Println(sex)

	// 同时赋予多个变量值
	var n1, n2, name = 100, 1000, "xuan"
	fmt.Println(n1)
	fmt.Println(n2)
	fmt.Println(name)

	fmt.Println(username)
	fmt.Println(password)

	// 浮点数可以用科学计数法表示
	var fnum = 314e+2
	fmt.Println(fnum)
	// 浮点数可能会有精度的损失，通常情况下，建议使用：float64
	// go中默认的浮点类型就是：float64

}

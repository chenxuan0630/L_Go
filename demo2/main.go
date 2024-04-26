package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 基本数据类型：
	// 整数、浮点、字符、布尔、字符串

	// 定义字符类型的数据：
	var c1 byte = 'a'
	fmt.Println(c1)

	// 字母，数字，标点底层是按照ASCII码进行存储
	var c2 byte = 'a'
	fmt.Println(c2) // 97

	// 汉字，底层对应的是Unicode码值
	// 可以用int存储

	// 总结：Golang的字符对应的使用的是UTF-8编码
	// Unicode是对应的字符串，UTF-8是Unicode的其中的一种编码方案

	// 想显示对应的字符，必须采用格式化输出
	var c5 byte = 'A'
	fmt.Printf("c5对应的具体的字符为：%c", c5)

	// 转义字符
	// \r 光标回到本行的开头，后续输入会替换原有的字符
	fmt.Println("aaaaa\rbbb")

	// \t 制表符，八个字符凑一个制表符，缺字符的时候用空格补上
	fmt.Println("aaa\tbbbbbbbb")

	// 布尔类型bool
	// 布尔类型的数据只允许true和false
	var flag bool = true
	fmt.Println(flag)

	// 字符串是不可变的，字符串一旦定义好，其中字符的值不能改变
	var s string = "Just string"
	fmt.Println(s)

	// 字符串中有特殊字符，字符串的表示形式用反引号，可以保留原格式
	var longa string = `
aaefjaofiasdfa
fsv0p awqa94r-[q923v93t8q[-4t-]2[ur] T3 3tfgt	 ]
QAEEEE09U-[0239]`
	fmt.Println(longa)

	// go在不同类型的变量之间赋值时只能够显式的进行转换!!!

	// 使用fmt.Sprintf("%d",n1)
	var n int = 10
	var s1 string = fmt.Sprintf("%d", n)
	fmt.Println(s1)
	//fmt.Printf("s1对应的类型是：%T , s1 = %v ", s1, s1)

	// 导入"strconv"包，将string类型转换为基本数据类型
	// ParseInt这个函数的返回值有两个：（value int， err error）
	// 我们只关注得到的数据，err可以使用 _ 忽略
	var sss string = "3.14"
	var nums int64
	nums, _ = strconv.ParseInt(sss, 10, 64)
	fmt.Println(nums)
}

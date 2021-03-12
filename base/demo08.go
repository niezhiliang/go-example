package main

import (
	"fmt"
	"strconv"
)

/*
字符串转基本数据类型
*/
func main() {

	//string转int64
	sn1 := "100"
	//参数1 ：待转换的字符串  参数2：转为几进制  参数3：转为int多少
	//返回值为一个元组  返回值1：转换的结果 返回值2：异常信息
	//这里跟Python有点像 不想要的值可以用_忽略掉
	n1, _ := strconv.ParseInt(sn1, 10, 64)
	fmt.Printf("n1的类型为：%t 值为：%v \n", n1, n1)

	//string转float34
	sf1 := "3.14"
	f1, _ := strconv.ParseFloat(sf1, 64)
	fmt.Printf("f1的类型为：%t 值为：%v \n", f1, f1)

	//string转bool
	sb1 := "true"
	b1, _ := strconv.ParseBool(sb1)
	fmt.Printf("b1的类型为：%t 值为：%v \n", b1, b1)

	//不是对应的string类型转换为bool，会转换为对应类型默认值
	sb2 := "GoLang"
	b2, _ := strconv.ParseBool(sb2)
	fmt.Printf("b1的类型为：%t 值为：%v \n", b2, b2)
}

package main

import (
	"fmt"
	"unsafe"
)

func main() {

	//值越界 无符号整数类型不能是负数
	//异常：constant -20 overflows uint8
	var num1 uint8 = -20
	fmt.Println("num1 = ", num1)

	//测试隐式获取变量的类型 和所占字节数
	num := 1
	//获取变量占用字节数
	size := unsafe.Sizeof(num)
	//格式化输出  %t获取变量的类型   %v获取变量的值
	fmt.Printf("num类型：%t 占用字节数量：%v", num, size)
}

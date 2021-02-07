package main

import (
	"fmt"
	"strconv"
	"unsafe"
)

func main() {
	//定义一个浮点类型变量
	num := 3.14
	//获取变量占用字节数
	size := unsafe.Sizeof(num)
	fmt.Printf("num数据类型：%t 占用字节数：%v",num,size)

	fmt.Println("\n<----------测试精度丢失------------>")
	//浮点转换为string类型，保留两位小数
	convertf := fmt.Sprintf("%.2f", 19.90)
	fmt.Printf("convertf类型：%t 值：%v \n",convertf,convertf)
	//string类型转换为float64
	num2, _ := strconv.ParseFloat(convertf, 64)
	//输出转换后的值
	fmt.Println("num2 = ",num2)
	//将该值乘于100
	fmt.Println("num2 * 100 = ",num2 * 100)

	fmt.Println("\n<----------精度不丢失方法------------>")
	/*
	num, _ := strconv.ParseFloat(fmt.Sprintf("%.8f", 19.90), 64)
	fmt.Println(num)

	decimalValue := decimal.NewFromFloat(num)
	decimalValue = decimalValue.Mul(decimal.NewFromInt(100))

	res,_ := decimalValue.Float64()
	fmt.Println(res)
	*/

}

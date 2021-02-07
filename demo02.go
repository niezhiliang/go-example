package main

import (
	"fmt"
)

//全局变量（第一种写法）
var a1 = 33
var b1 = 44

//全局变量（第二种写法）
var (
	c1 = 29
	d1 = 92
)

//定义全局常量
const (
	e1 = 22
	f1 = 33
)


func main()  {
	//声明变量的四种方法
	fmt.Println("<-------------局部变量定义演示------------->")
	//1.声明变量a（int）
	var a int
	//2.变量的赋值
	a = 18
	//3.变量的使用 输出到控制台
	fmt.Println("a = ",a)

	//1声明 2、赋值 合为一句代码（定义变量b=20）
	var b int = 20
	//3.变量的使用 输出到控制台
	fmt.Println("b = ",b)

	//1声明 2、赋值 合为一句代码 类型省略 隐式推断类型 定义变量c=20）
	var c  = 22
	//3.变量的使用 输出到控制台
	fmt.Println("c = ",c)

	//等同于上面这种写法
	d := 25
	//变量的使用 输出到控制台
	fmt.Println("d = ",d)

	//变量赋值不对应的值 报异常.\demo02.go:26:6: constant 2.5 truncated to integer
	//var e int = 3.14
	//fmt.Println("e = ",e)
	fmt.Println("<-------------全局变量定义演示------------->")
	//打印全局变量的值到控制台
	fmt.Printf("a1 = %v b1 = %v c1 = %v d1 = %v\n",a1,b1,c1,d1)

	fmt.Println("<-------------常量定义演示------------->")
	//常量定义
	const e int = 66
	//隐式推断类型
	const f ="常量字符串"

	//错误示例：常量不允许更改
	//e = 77

	//将常量输出到控制台
	fmt.Printf("e = %v f = %v e1 = %v f1 = %v",e,f,e1,f1)
}
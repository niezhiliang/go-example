package main

import "fmt"

/*
  修改变量的值
*/
func test(age int) {
	age = 20
	fmt.Println("test函数age = ", age)
}

/*
  修改变量的值，通过指针改变变量的值
*/
func test2(age *int) {
	*age = 25
	fmt.Println("test2函数age = ", age)
}

/*
接收函数参数的函数
*/
func test3(age int, test func(int)) {
	test(age)
}

/*
程序入口
*/
func main() {
	//函数赋值给变量
	a := test
	//定义年龄变量
	age := 18
	//变量也可以当函数调用
	a(age)
	fmt.Println("main函数中的age = ", age)
	//通过指针改变age的值
	test2(&age)
	fmt.Println("main函数中的age = ", age)
	//调用函数有形参为函数的函数
	test3(18, test)
}

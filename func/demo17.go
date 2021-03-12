package main

import "fmt"

func main() {
	//调用无返回值
	add1(1, 5)
	//一个返回值对象
	rResult := reduce(10, 3)
	fmt.Println("10 - 5 = ", rResult)
	//多个返回值对象
	addResult, reduceResult := addAndReduce(20, 10)
	fmt.Println("20 + 10 = ", addResult, "\n20 - 10 = ", reduceResult)
	//调用取了别名的函数  只关心加法，用_忽略减法
	result, _ := addAndReduce2(30, 20)
	fmt.Println("30 + 20 = ", result)
}

/*
  两数相加
  无返回值函数
*/
func add1(num1 int, num2 int) {
	fmt.Printf("%v + %v = %v\n", num1, num2, num1+num2)
}

/*
 数相减
 返回一个值
*/
func reduce(num1 int, num2 int) int {
	return num1 - num2
}

/**
  同时做加法和减法
  有返回多个值函数
*/
func addAndReduce(num1 int, num2 int) (int, int) {
	sum := num1 + num2
	result := num1 - num2
	return sum, result
}

/**
  同时做加法和减法，返回值取别名
  有返回多个值函数
*/
func addAndReduce2(num1 int, num2 int) (sum int, result int) {
	sum = num1 + num2
	result = num1 - num2
	return
}

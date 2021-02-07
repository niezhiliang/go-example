package main

import "fmt"

/*
  数组转切片
*/
func main() {
	array := [10] int {1,2,3,4,5,6,7,8,9,10}
	//数组转对象，取数组中所有元素
	slice1 := array[:]
	fmt.Println("slice1:",slice1)
	//数组转对象，从第一个开始取
	slice2 := array[1:]
	fmt.Println("slice2:",slice2)
	//数组转对象，从第一个开始取到 5 - 1 个
	slice3 := array[:5]
	fmt.Println("slice3:",slice3)
	//修改数组最后一个元素为100
	array[9] = 100
	fmt.Println("slice1:",slice1)
	//slice1 = append(slice1, 11)
	//array[8] = 120
	//fmt.Println("slice1:",slice1)
	//创建一个切片
	slice4 := make([]int, 10)
	//将切片slice1元素拷贝给slice4
	copy(slice4,slice1)
	//修改老数组的值
	array[8] = 110
	fmt.Println("slice1:",slice1)
	fmt.Println("slice4:",slice4)

}

package main

import "fmt"

func main()  {

	//定义一个空数组,默认填充对应类型的默认值
	var array01 [10] int
	//给数组第一个元素赋值
	array01[0] = 2
	fmt.Println("array01:",array01)

	//定义一个长度为5的数组，并给定值
	array02 := [5] int {1,2,3,4,5}
	fmt.Println("array02:",array02)

	//定义一个未知长度的数组，编译器会根据元素推断出数组长度
	//编译完成后，假设现在是10个元素，后期在方法中也不能给该数组加元素
	//你可以理解为编译时，go帮你把...替换成了10就行
	array03 := [...] int {1,3,5,6,9,7,5,2,2,6}
	fmt.Printf("array中元素个数 = %v \n",len(array03))

}

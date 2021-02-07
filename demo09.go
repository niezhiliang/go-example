package main

import "fmt"

func main() {
	//定义一个int类型的变量
	var age int = 18
	////通过&可以获取到变量的指针，&表示取地址符
	////指针必须使用对应的指针变量去接收 指针变量就是在原来的数据类型前面加了个*
	//var ptr *int = &age
	//fmt.Printf("age的指针为：%v,值为：%v \n",ptr , *ptr)
	//
	////改变指针对应变量的值
	//*ptr = 25
	//fmt.Printf("指针修改后age的值：%v \n",age)

	//调用a函数
	a(age)
	fmt.Printf("函数修改后age的值：%v \n",age)

}

func a(age int)  {
	age = 20
	fmt.Println("a函数内的age值：",age)
}

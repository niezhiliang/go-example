package main

import "fmt"

func main()  {
	//（方式一）定义一个空切片,默认值为 nil
	var slice1 [] int
	fmt.Println("slice1",slice1)

	//（方式二）通过make方法定义一个长度为0容量为3的切片,超过3自动扩容
	slice2 := make([]int,0,3)
	fmt.Println("slice2",slice2)
	//区别 slice1=nil slice2相当于初始化了，不为nil

	//定义带初始值的切片，支持自动扩容
	slice3 := [] int {1,2,3,4}
	fmt.Println("slice3：",slice3)

	//给slice3添加元素
	slice3 = append(slice3, 1)
	fmt.Println("slice3：",slice3)
	//删除slice3的第一个元素
	slice3 = remove(slice3, 0)
	fmt.Println("slice3：",slice3)

	lenth := len(slice3)
	capsize := cap(slice3)
	fmt.Println("切片的长度：",lenth," 容量：",capsize)

}

/*
移除下标为index的元素
*/
func remove(slice [] int,index int) [] int {
	return append(slice[:index],slice[index + 1:]...)
}

package main

import (
	"fmt"
	"time"
)

var ch = make(chan int)

var ch2 = make(chan int,2)

func main()  {
	//var num1 float64 = 5.3611
	//var num2 float64 = 0.5311
	//res := num1 * num2
	//fmt.Println()
	//num := 10
	//fmt.Println("num = v%",num)
	//numPoint := &num
	//fmt.Println("numPoint = v%",num)
	//*numPoint = 20
	//fmt.Println("num = ",num)
	//
	//fmt.Println("hello go")
	//now := utils.GetDateNow()
	//fmt.Println(now)
	//result,result2 := add(10, 20)
	//fmt.Println(result,result2)
	//
	//fmt.Printf("num1的类型：%t",num1)
	//
	//if res >10 {
	//	fmt.Println("res的值小于10")
	//} else {
	//	fmt.Println("res的值大于10")
	//}
	//
	//str := "25"
	////字符串转换其他类型
	//parseInt, _ := strconv.ParseInt(str, 10, 8)
	//
	//fmt.Printf("parseInt的值为：%v 类型为:%t,",parseInt,parseInt)
	//
	//strResult := strconv.Itoa(result)
	//fmt.Println("-----------------")
	//fmt.Printf("strResult值为：%v, 类型为：%t",strResult,strResult )
	//fmt.Println("-----------------")
	////声明数组
	//array := [5] int {1,2,3,4,5}
	//
	//for i := range array {
	//	fmt.Println(array[i])
	//}
	////切片
	//slice := make([] int, 0)
	//slice = append(slice, 1)
	//slice = append(slice, 2)
	//
	//fmt.Println("切片长度：" , len(slice))
	//
	//for i,v := range slice {
	//	fmt.Printf("i：%v  v:%v  \n" , i,v)
	//}
	//冒泡排序
	array := [] int {6,8,2,4,7,5,3}

	for i := 0; i < len(array) -1; i++ {
		for j := 0; j < len(array) - 1 -i; j++ {
			if array[j] > array[j + 1] {
				temp := array [j]
				array[j] = array[j + 1]
				array[j + 1] = temp
			}
		}
	}

	for _, value := range array {
		fmt.Println(value)
	}

	ints := make([]int ,5)

	ints = append(ints, 1)

	for i := range ints {
		fmt.Println(ints[i])
	}

	fruit := make(map[string]string,10)
	fruit["apple"] = "苹果"
	fruit["bananer"] = "香蕉"

	fmt.Printf("%v \n",fruit)

	for key, value := range fruit {
		fmt.Println(key,value)
	}

	apple := fruit["apple"]

	fmt.Printf("apple value: %v \n", apple)
	delete(fruit, "apple")
	fmt.Printf("%v \n",fruit)


	//创建一个协程
	//go myTask()
	//go person1()
	//go person2()

	go func() {
		for i := 0; i < 4; i++ {
			fmt.Println("i=",i)
			ch2 <- i
			fmt.Printf("len(ch2) = %d, cap(ch2) = %d, cap(ch2) = %d \n",i, len(ch2), cap(ch2))
		}
	}()

	go func() {
		for i := 0; i < 4; i++ {
			num := <- ch2
			fmt.Println("num:",num)
		}
	}()

	for {

	}

}

/*
两个数相加
 */
func add(num1 int,num2 int) (sum int,dev int)  {
	sum = num1 + num2
	dev = num1 - num2
	return
}

func myTask()  {
	for {
		fmt.Println("hello ",time.Now())
		time.Sleep(time.Second)
	}
}

func printer(str string)  {
	for _, vale := range str {
		fmt.Printf("%c",vale)
		time.Sleep(time.Second)
	}
	fmt.Println("\n")
}

func person1()  {
	printer("hello")
	ch <- 666 //给管道写数据 发送
}
func person2()  {
	<- ch //从管道读取数据，没有数据会阻塞
	printer("world")
}

func m(f func(str string))  {
	//f()
}



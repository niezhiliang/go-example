package main

import "fmt"

func main() {
	//类似于java里面的while
	i := 0 //变量的初始化
	for i < 5 { //条件表达式,判断条件
		fmt.Println(i)
		i++//累加
	}

	//for range 有点像java里面的foreach
	str := "hello"
	//遍历的是索引和值
	for index, value := range str {
		fmt.Printf("索引:%v, value:%c \n",index,value)
	}
	//给外层循环设置一个label
	label:
	for i := 0; i < 5; i++ {
		for j := 2; j < 4; j++ {
			fmt.Printf("i:%v, j:%v \n",i,j)
			if i == 2 && j == 2 {
				break label //结束指定标签对应的循环，如果标签定义了，不用会报错
			}
		}
	}

	label1:
	for i := 0; i < 3; i++ {
		for j := 2; j < 4; j++ {
			fmt.Printf("i:%v, j:%v \n",i,j)
			if i == 2 && j == 2 {
				continue label1 //结束指定标签对应的循环，如果标签定义了，不用会报错
			}
		}
	}
	fmt.Println("------ok------")


	fmt.Println("开始执行goto关键词")
	goto label2
	fmt.Println("这里的内容将直接被跳过")
	label2:
		fmt.Println("goto跳转过来后执行的代码")
}

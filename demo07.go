package main

import (
	"fmt"
	"strconv"
)

/*
 基本数据类型转string
*/
func main() {
	//int类型转string
	n1 := 110
	//方式一（推荐方式）
	sn1 := fmt.Sprintf("%d",n1)
	fmt.Printf("sn1数据类型：%t 值为：%v \n",sn1,sn1)

	//方式二
	sn2 := strconv.Itoa(n1)
	fmt.Printf("sn2数据类型：%t 值为：%v \n",sn2,sn2)

	//float64转string
	f1 := 3.14
	//方式一（推荐）
	sf1 := fmt.Sprintf("%.3f",f1)
	fmt.Printf("sf1数据类型：%t 值为：%v \n",sf1,sf1)
	//第二个参数：'f'（-ddd.dddd）  第三个参数：9 保留小数点后面9位  第四个参数：表示这个小数是float64类型
	sf2 := strconv.FormatFloat(f1,'f',3,64)
	fmt.Printf("sf2数据类型：%t 值为：%v \n",sf2,sf2)

	//bool转string
	var flag bool
	//方式一（推荐）
	sb1 := fmt.Sprintf("%d", flag)
	fmt.Printf("sb1数据类型：%t 值为：%v \n",sb1,sb1)
	//方式二
	sb2 := strconv.FormatBool(flag)
	fmt.Printf("sb2数据类型：%t 值为：%v \n",sb2,sb2)

	//字符类型转string
	c := 'a'
	//第一种方式
	sc1 := fmt.Sprintf("%c", c)
	fmt.Printf("sc1数据类型：%t 值为：%v \n",sc1,sc1)
}

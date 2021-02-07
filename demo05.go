package main

import "fmt"

func main() {
	//定义字符类型变量 值为a
	var c1 byte = 'a'
	//输出a对应的ASCII码和输出该ASCII码对应的值
	fmt.Printf("c1对应的ASCII码为：%v, 对应的字符为：%c \n",c1,c1)

	c2 := '6'
	fmt.Printf("c2对应的ASCII码为：%v, 对应的字符为：%c \n",c2,c2)

	c3 := '('
	fmt.Printf("c3对应的ASCII码为：%v, 对应的字符为：%c \n",c3,c3)

	//中字为汉字，底层对应Unicode码，对应20013，超过了byte范围的值，所有会用int来接收
	c4 := '中'
	fmt.Printf("c4对应的ASCII码为：%v, 对应的字符为：%c \n",c4,c4)

}

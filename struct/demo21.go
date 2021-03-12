/**
 * @Author: niezhiliang
 * @Date: 2021/3/12 11:13 上午
 */

package main

import "fmt"

//定义数据结构体
type MyBook struct {
	//名称
	name string
	//作者
	author string
	//学科
	subject string
	//编号
	bookId int
}

//为结构体添加方法
func (book MyBook) say() {
	fmt.Println("book:", book)
}

type MyStudent struct {
	//继承MyBook,属性和函数都继承了，如果属性一样，继承的类会覆盖父类的属性
	MyBook        //不需要变量名，写了变量名是引用该结构体
	b      MyBook //引用了MyBook这个结构体
	//姓名
	name string
	//年龄
	age int
}

//为结构体添加方法
/*
func (student MyStudent) say()  {
	fmt.Println("student:",student)
}
*/

func main() {
	b := MyBook{"Java", "詹姆斯搞死你", "计算机", 1}
	a := MyStudent{b, b, "张三", 18}
	//直接调用MyBook中的say函数
	a.say()
	fmt.Println(a)
}

/**
 * @Author: niezhiliang
 * @Date: 2021/2/9 4:58 下午
 */

package main

import "fmt"

//定义数据结构体
type Book struct {
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
func (book Book) say() {
	fmt.Println("book:", book)
}

/*
结构体作为参数传给函数
*/
func printBook(book Book) {
	book.name = "haha"
	fmt.Println("函数内book2:", book)
}

/*
函数接受结构体的指针
*/
func printBook2(book *Book) {
	book.name = "haha"
	fmt.Println("函数内book2:", *book)
}
func main() {
	//创建一个Book的结构体，类似Java中的构造创建对象
	book := Book{"Go", "张三", "编程", 10086}
	//调用结构体的方法
	book.say()
	fmt.Println("book:", book)
	//创建结构体第二种方式，没有赋值的属性默认为数据类型对应的默认值
	book2 := Book{name: "Java", author: "李四"}
	//结构体作为参数传给函数，传递的是结构体的值
	printBook(book2)
	fmt.Println("book2:", book2)
	//传递指针给函数
	printBook2(&book2)
	fmt.Println("book2:", book2)
	book3 := Book{}
	//设置name属性的值
	book3.name = "C"
	book3.author = "王五"
	book3.subject = "编程"

	fmt.Println("book3.name:", book3.name, " author:", book3.author, " subject:", book3.subject)

}

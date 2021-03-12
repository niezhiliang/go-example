/**
 * @Author: niezhiliang
 * @Date: 2021/3/12 1:13 下午
 */

package main

import (
	"fmt"
)

//定义一个结构体Student
type Student struct {
	//定义name属性
	name string
}

//定义一个结构体Teacher
type Teacher struct {
	//定义name属性
	name string
}

//为结构体Student实现interface Person的say函数
func (s Student) Say() {
	fmt.Println("student", s.name, " say hello")
}

//为结构体Teacher实现interface Person的say函数
func (t Teacher) Say() {
	fmt.Println("teacher ", t.name, " say hello")
}

func WhoSay(p Person) {
	p.Say()
}

func main() {
	//用interface来接收结构体的变量
	var s Person = Student{"张三"}
	t := Teacher{"李四"}
	//不同表现 多态的体现
	s.Say()
	t.Say()

	WhoSay(s)
	WhoSay(t)

}

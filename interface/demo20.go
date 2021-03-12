/**
 * @Author: niezhiliang
 * @Date: 2021/2/9 7:14 下午
 */

package main

//type Person interface {
//	RealName() string
//}
//
//type Role interface {
//	Person //接口继承
//	ReturnRole() string
//}
//
//type Student struct {
//	name string
//}
//
//func (s *Student) RealName() string {
//	return s.name
//}
//
//func (s Student) ReturnRole() string {
//	return "学生"
//}
//
//type Worker struct {
//	name string
//}
//
//func (w *Worker) RealName() string {
//	return w.name
//}
//
//func main() {
//	student := Student{"张三"}
//	var person Person
//	person = &student
//	name := person.RealName()
//	fmt.Println(name)
//
//	worker := Worker{"哈哈"}
//	person2 := &worker
//	name2 := person2.RealName()
//	fmt.Println(name2)
//
//	CheckPerson(student)
//
//	//s := Student{"李四"}
//	//var role Role
//	//role = s
//
//}
//
//func CheckPerson(test interface{}) {
//	if _, ok := test.(Person); ok {
//		fmt.Println("Student implement Person")
//	}
//}

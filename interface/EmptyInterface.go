/**
 * @Author: niezhiliang
 * @Date: 2021/3/12 2:58 下午
 */

package main

import "fmt"

type MyStruct struct {
	name string
	age  int
}

//返回值为空接口
func testReturnEmptyInterface() interface{} {
	return MyStruct{"张三", 20}
}

func main() {
	//空接口，可以用来接收任何类型
	var i interface{} = 5
	fmt.Println(i)
	emptyInterface := testReturnEmptyInterface()
	fmt.Printf("%t \n", emptyInterface)
	//强转
	var structs = emptyInterface.(MyStruct)
	fmt.Println(structs)

}

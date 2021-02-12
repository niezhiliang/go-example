package main

import "fmt"

func main() {
	//定义一个map变量 初始值为nil，没有初始化不能往集合中添加元素
	var map1 map[string]int
	fmt.Println("map1", map1)
	//异常：panic: assignment to entry in nil map
	//map1["key"] = 1
	//定义一个带初始值的map
	map2 := map[string]int{"a": 1, "b": 2}
	fmt.Println("map2", map2)

	//定义一个初始化的空map
	map3 := make(map[string]int, 10)
	fmt.Println("map3", map3)
	//map2添加元素 key = a  value = 1
	map3["c"] = 1
	map3["d"] = 2
	fmt.Println("map3", map3)

	//通过key获取value 会返回一个元组
	//第一个元素为key对应的值，第二个表示是否存在
	value, ok := map3["d"]
	fmt.Println("value:", value, "ok", ok)

	//删除map3中的元素
	delete(map3, "d")
	fmt.Println("map3", map3)
}

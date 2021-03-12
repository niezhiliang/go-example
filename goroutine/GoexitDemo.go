/**
 * @Author: niezhiliang
 * @Date: 2021/3/12 10:18 上午
 */

package main

import (
	"fmt"
	"runtime"
	"time"
)

/*
不加Goexit执行效果：
B
B.defer
A
A.defer
加上Goexit效果
B.defer
A.defer
*/
func main() {
	go func() {
		//最后执行 类似java中try finally中的代码
		defer fmt.Println("A.defer")
		//匿名函数
		func() {
			defer fmt.Println("B.defer")
			runtime.Goexit()
			fmt.Println("B")
		}()
		fmt.Println("A")
	}()

	time.Sleep(5 * time.Second)
}

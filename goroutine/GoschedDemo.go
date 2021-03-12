/**
 * @Author: niezhiliang
 * @Date: 2021/3/12 10:05 上午
 */

package main

import (
	"fmt"
	"time"
)

/*
	注释 runtime.Gosched() 输出效果为：
	i am main routine
	i am main routine
	i am a goroutine
	i am a goroutine

	打开后输出效果如下：

	i am a goroutine
	i am main routine
	i am a goroutine
	i am main routine
*/
func main() {
	//创建一个携程
	go func(s string) {
		for i := 0; i < 2; i++ {
			fmt.Println(s)
		}
	}("i am a goroutine")

	for i := 0; i < 2; i++ {
		//让出cpu的时间片，
		//runtime.Gosched()
		fmt.Println("i am main routine")
	}
	//防止主携程退出
	time.Sleep(time.Second)
}

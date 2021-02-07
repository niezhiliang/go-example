package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	flag := true
	//小括号可以省略，变量和if关键词之间一定要有空格
	//{}必须写，就算只有一行代码也不能省略
	//单分支
	if flag {
		fmt.Println("flag：",flag)
	}

	//加入种子资源库，不然每次生成的随机数都是一样的
	//因为rand是个伪随机生成器，而同一个共享源生成的随机数序列是恒定的。
	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(20)
	//双分支
	if num < 10 {
		fmt.Printf("随机数【%v】小于10",num)
	} else {
		fmt.Printf("随机数【%v】大于10",num)
	}

	//多分支
	if num < 5 {
		fmt.Printf("随机数【%v】小于5",num)
	} else if num < 10 {
		fmt.Printf("随机数【%v】大于5小于10",num)
	} else {
		fmt.Printf("随机数【%v】大于10",num)
	}
}

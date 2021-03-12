package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())
	//随机一个分数
	score := rand.Intn(101)
	fmt.Println("你的分数为：", score)
	//分数除以10 得到评分
	//结果10、9 评分为A  8、7 评分为B  6、5评分为C 4、3评分为D 小于3的全齐退学吧
	switch score / 10 {
	case 10, 9:
		fmt.Println("你的等级为A")
		//可以省略，后面跟fallthrough 也不执行
		break
	case 8:
		//穿透到下一个case里面，只能穿一层
		fallthrough
	case 7:
		fmt.Println("你的等级为B")
	case 6:
		fallthrough
	case 5:
		fmt.Println("你的等级为C")
	//default随便写哪里
	default:
		fmt.Println("你没救了，退学吧")
	case 4, 3:
		fmt.Println("你的等级为G")
	}
}

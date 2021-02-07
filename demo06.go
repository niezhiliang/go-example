package main

import "fmt"

func main()  {
	fmt.Println("<--------int转浮点---------------->")

	n1 := 100
	f2 :=float32(n1)

	fmt.Println("int类型的100转为float32，结果：",f2)

	fmt.Println("<--------浮点转int---------------->")

	f3 := 3.14
	var n2 = int(f3)

	fmt.Println("3.14转为int，结果：",n2)

	fmt.Println("<--------int8转int---------------->")
	var n3 int8 = 125
	var n4 = int(n3)

	fmt.Println("int8类型的125转int64，结果：",n4)

	fmt.Println("<--------int转int8---------------->")
	//这里比较好玩，如果是128 转出来的是-128，129-> -127 依次类推   256为0 如果再大就会是 1，有点一致性哈希环的感觉，不管你传啥都会在-128 - 127之间
	//所以在用的时候要特别小心
	n5  := 128
	var n6 = int8(n5)
	fmt.Println("int类型的128转int8，结果：",n6)

	fmt.Println("<--------int转uint---------------->")
	//这里跟上面一样，会取到正值
	var n7 int = -300
	var n8 = uint(n7)

	fmt.Println("int类型的-300转uint，结果：",n8)
}

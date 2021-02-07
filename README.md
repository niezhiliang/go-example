### 第一个Go程序

```go
package main //demo01

//导包
import "fmt"

/*
  程序入口（多行注释）
 */
func main() {
	//第一个go程序（单行注释）
	fmt.Println("hello GoLand")
}
```

### 语法注意事项

1. 源文件以"go"为扩展名。
2. 程序的执行入口是main()函数。
3. 严格区分大小写。
4. 方法由一条条语句构成，每个语句后不需要分号(Go语言会在每行后自动加分号)，这也体现出Golang的简洁性。
5. Go编译器是一行行进行编译的，因此我们一行就写一条语句，不能把多条语句写在同一个，否则报错
6. **定义的变量或者import的包如果没有使用到，代码不能编译通过**。
7. 大括号都是成对出现的，缺一不可

### API文档

Golang中文网在线标准库文档: https://studygolang.com/pkgdoc

### 变量

**变量使用的步骤和声明变量的四种方法**

1. 声明
2. 赋值
3. 使用

```go
//1.声明变量a（int）//demo02
var a int
//2.变量的赋值
a = 18
//3.变量的使用
fmt.Println("a = ",a)

//1声明 2、赋值 合为一句代码（定义变量b=20）
var b int = 20
//3.变量的使用
fmt.Println("b = ",b)

//1声明 2、赋值 合为一句代码 类型省略 会自动获取定义变量c=20）
var c  = 22
//3.变量的使用
fmt.Println("c = ",c)

//等同于上面这种写法
d := 25
//变量的使用
fmt.Println("d = ",d)
```

**注意：** 不可以在赋值的时候给与不匹配的类型，go编译不会报错，运行时会报错

```go
//定义一个in类型变量、赋值给了一个浮点类型
var i int = 2.5
fmt.Println(i)
```

上面代码会得到下面的异常

```shell
# command-line-arguments
.\demo02.go:26:6: constant 2.5 truncated to integer
```

**全局变量的定义**

```go
package main //demo02

import (
	"fmt"
)

//全局变量（第一种写法）
var a1 = 33
var b1 = 44

//全局变量（第二种写法）
var (
	c1 = 29
	d1 = 92
)


func main()  {
	//打印全局变量的值到控制台
	fmt.Printf("a1 = %v b1 = %v c1 = %v d1 = %v",a1,b1,c1,d1)
}
```

**常量定义**

定义语法：const  常量名  常量类型 = 常量值

```go
package main

import (
	"fmt"
)

//定义全局常量
const (
	e1 = 22
	f1 = 33
)

func main()  {
	//常量定义
	const e int = 66
	//隐式推断类型
	const f ="常量字符串"
    
	//错误示例：常量不允许更改,编译都过不去
	//e = 77

	//将常量输出到控制台
	fmt.Printf("e = %v f = %v",e,f)
}
```

### 数据类型

Go的数据类型分为了两种，基本数据类型和复杂数据类型

#### 基本数据类型

整数型、浮点型、字符类型、布尔类型、字符串类型

##### 整数型

整数型又分为了有符号整数类型和无符号整数类型以及其他整数类型，有符号整数类型能有负数，无符号类型的值只能大于0。

- 有符号整数类型：`int8`、`int16`、`int32`、`int64` 分别对应java中的`byte`、`short`、`int`、`long`

![](https://java-imgs.oss-cn-hongkong.aliyuncs.com/2021/1/8/image_41538.png)

- 无符号整数类型：没有负数

![](https://java-imgs.oss-cn-hongkong.aliyuncs.com/2021/1/8/image_30207.png)

- 其他整数类型

![](https://java-imgs.oss-cn-hongkong.aliyuncs.com/2021/1/8/image_51532.png)

**友情提醒**

1.定义好的类型赋值不能越界

```go
// uint为无符号整数类型，不支持负数，这里肯定会报异常 //demo03
var num1 uint8 = -20
fmt.Println("num1 = ",num1)
```

```shell
constant -20 overflows uint8
```

2.隐式获取整数类型的变量默认为`int`类型，64位系统占8个字节

```go
//测试隐式获取变量的类型 和所占字节数  //demo03
num := 1
//获取变量占用字节数
size := unsafe.Sizeof(num)
//格式化输出  %t获取变量的类型   %v获取变量的值
fmt.Printf("num类型：%t 占用字节数量：%v",num,size)
```

```shell
num类型：%!t(int=1) 占用字节数量：8
```

3.这么多整数类型，使用的时候该如何选择呢？**在保证程序正确运行下,尽量使用占用空间小的数据类型**

4.变量只声明不赋值的情况默认值为0

##### 浮点型

浮点类型分为了`float32`和`float64`

![](https://java-imgs.oss-cn-hongkong.aliyuncs.com/2021/1/8/image_54693.png)

**友情提醒**

1.隐式获取浮点类型的变量默认为`float64`类型，占8个字节，变量不赋值默认为0

```go
//定义一个浮点类型变量 //demo04
num := 3.14
//获取变量占用字节数
size := unsafe.Sizeof(num)
fmt.Printf("num数据类型：%t 占用字节数：%v",num,size)
```

```shell
num数据类型：%!t(float64=3.14) 占用字节数：8
```

2.不管是`float32`还是`float64`都会丢失精度，建议使用`float64`、实在不行就用`decimal`,不过`decimal`并不是`go`sdk中提供的

```go
//浮点转换为string类型，保留两位小数 //demo04
convertf := fmt.Sprintf("%.2f", 19.90)
fmt.Printf("convertf类型：%t 值：%v \n",convertf,convertf)
//string类型转换为float64
num2, _ := strconv.ParseFloat(convertf, 64)
//输出转换后的值
fmt.Println("num2 = ",num2)
//将该值乘于100
fmt.Println("num2 * 100 = ",num2 * 100)
```

```shell
convertf类型：%!t(string=19.90) 值：19.90 
num2 =  19.9
num2 * 100 =  1989.9999999999998
```

//需要引入第三方包，放后面再说 github.com/shopspring/decimal

```go
//将字符串转换成float64  //demo04
num, _ := strconv.ParseFloat(fmt.Sprintf("%.8f", 19.90), 64)
fmt.Println(num)

decimalValue := decimal.NewFromFloat(num)
decimalValue = decimalValue.Mul(decimal.NewFromInt(100))

res,_ := decimalValue.Float64()
fmt.Println(res
```

##### 字符类型

go中没有专门的字符类型，如果要存储单个字符（字母），一般使用byte保存，go默认使用的字符为UTF-8

```go
//定义字符类型变量 值为a  //demo05
var c1 byte = 'a'
//输出a对应的ASCII码和输出该ASCII码对应的值
fmt.Printf("c1对应的ASCII码为：%v, 对应的字符为：%c \n",c1,c1)

c2 := '6'
fmt.Printf("c2对应的ASCII码为：%v, 对应的字符为：%c \n",c2,c2)

c3 := '('
fmt.Printf("c3对应的ASCII码为：%v, 对应的字符为：%c \n",c3,c3)

//中字为汉字，底层对应Unicode码，对应20013，超过了byte范围的值，所有会用int来接收
c4 := '中'
fmt.Printf("c4对应的ASCII码为：%v, 对应的字符为：%c \n",c4,c4)
```

```shell
c1对应的ASCII码为：97, 对应的字符为：a 
c2对应的ASCII码为：54, 对应的字符为：6 
c3对应的ASCII码为：40, 对应的字符为：( 
c4对应的ASCII码为：20013, 对应的字符为：中 
```

该类型也没啥纠结的，实际开发中基本不用该类型，完全可以用string来代替

##### 布尔类型

关键词为`bool` ,占用1个字节，等价于Java中的`boolean`,**不赋值默认为false**

##### 字符串

go中字符串也被归为基本的数据类型，关键词为`string`用法跟Java中大同小异，

**友情提醒**

1.如果字符串太长，需要换行的话，**使用+号拼接，不过+号必须留在上一行**，不然编译器会认为上一行代码还没有写完

![](https://java-imgs.oss-cn-hongkong.aliyuncs.com/2021/1/8/20210206174358.png)

2.声明一个string类型的字符串，不给默认值的，默认为“”

```go
var s1 string
var s2 string = ""
flag := strings.EqualFold(s1,s2)
fmt.Println("判断s1等于s2,结果：",flag)
```

```shell
判断s1等于s2,结果： true
```

##### 数据类型之间的转换

```go
//demo06
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

```

```shell
<--------int转浮点---------------->
int类型的100转为float32，结果： 100
<--------浮点转int---------------->
3.14转为int，结果： 3
<--------int8转int---------------->
int8类型的125转int64，结果： 125
<--------int转int8---------------->
int类型的128转int8，结果： -128
<--------int转uint---------------->
int类型的-300转uint，结果： 18446744073709551316
```



<img src="https://java-imgs.oss-cn-hongkong.aliyuncs.com/2021/1/8/20210206182722.png" style="zoom:80%;float:left" />

##### 基本数据类型转string

转换共有两种方式，推荐使用第一种方式

- 方式1：fmt.Sprintf(format string, a ...interface{}) ` api链接https://studygolang.com/pkgdoc 搜索 fmt`

  ![](https://java-imgs.oss-cn-hongkong.aliyuncs.com/2021/1/8/20210206190655.png)

```go
//int类型转string  //demo07
n1 := 110
//方式一（推荐方式）
sn1 := fmt.Sprintf("%d",n1)
fmt.Printf("sn1数据类型：%t 值为：%v \n",sn1,sn1)

//方式二
sn2 := strconv.Itoa(n1)
fmt.Printf("sn2数据类型：%t 值为：%v \n",sn2,sn2)

//float64转string
f1 := 3.14
//方式一（推荐）
sf1 := fmt.Sprintf("%.3f",f1)
fmt.Printf("sf1数据类型：%t 值为：%v \n",sf1,sf1)
//第二个参数：'f'（-ddd.dddd）  第三个参数：9 保留小数点后面9位  第四个参数：表示这个小数是float64类型
sf2 := strconv.FormatFloat(f1,'f',3,64)
fmt.Printf("sf2数据类型：%t 值为：%v \n",sf2,sf2)

//bool转string
var flag bool
//方式一（推荐）
sb1 := fmt.Sprintf("%d", flag)
fmt.Printf("sb1数据类型：%t 值为：%v \n",sb1,sb1)
//方式二
sb2 := strconv.FormatBool(flag)
fmt.Printf("sb2数据类型：%t 值为：%v \n",sb2,sb2)

//字符类型转string
c := 'a'
//第一种方式
sc1 := fmt.Sprintf("%c", c)
fmt.Printf("sc1数据类型：%t 值为：%v \n",sc1,sc1)
```

* 方式2：strconv.FormatXxxx()  api跟上面一样搜索 strcover

![](https://java-imgs.oss-cn-hongkong.aliyuncs.com/2021/1/8/20210206191532.png)

##### string类型转基本数据类型

```go
//string转int64 //demo08
sn1 := "100"
//参数1 ：待转换的字符串  参数2：转为几进制  参数3：转为int多少
//返回值为一个元组  返回值1：转换的结果 返回值2：异常信息
//这里跟Python有点像 不想要的值可以用_忽略掉
n1, _ := strconv.ParseInt(sn1, 10, 64)
fmt.Printf("n1的类型为：%t 值为：%v \n",n1,n1)

//string转float34
sf1 := "3.14"
f1, _ := strconv.ParseFloat(sf1, 64)
fmt.Printf("f1的类型为：%t 值为：%v \n",f1,f1)

//string转bool
sb1 := "true"
b1, _ := strconv.ParseBool(sb1)
fmt.Printf("b1的类型为：%t 值为：%v \n",b1,b1)

//不是对应的string类型转换为bool，会转换为对应类型默认值
sb2 := "GoLang"
b2, _ := strconv.ParseBool(sb2)
fmt.Printf("b1的类型为：%t 值为：%v \n",b2,b2)
```

**友情提醒**

如果string类型的值不是转换的目标类型，则会取对应类型的默认值

#### 派生数据类型/复杂数据类型

##### 指针

每个基本数据类型变量，内存都会为其开辟一块内存空间，这个空间的地址值，我们就叫它指针，看下图，大家应该会更加清楚指针的概念

![](https://java-imgs.oss-cn-hongkong.aliyuncs.com/2021/1/8/image_20936.png)

**取变量指针**

使用`&变量`就能得到变量的指针，示例：`&age`。变量的指针只能用`指针变量`来接收，

**指针变量**

指针变量的定义在原本的变量数据类型前加上`*`，例如 var ptr *int = &age 。

**指针修改变量**

变量也可以通过指针直接改变变量的值。示例 *ptr = 20，我们就将age的值改为了20

```go
//定义一个int类型的变量 //demo09
var age int = 18
//通过&可以获取到变量的指针，&表示取地址符
//指针必须使用对应的指针变量去接收 指针变量就是在原来的数据类型前面加了个*
var ptr *int = &age
fmt.Printf("age的指针为：%v,值为：%v \n",ptr , *ptr)

//改变指针对应变量的值
*ptr = 25
fmt.Printf("指针修改后age的值：%v \n",age)
```

```shell
age的指针为：0xc00000a0a0,值为：18 
指针修改后age的值：25 
a函数内的age值： 20
```

我们只需要记住，指针关键的两个符号，**&取内存地址，*根据地址取值**

**友情提醒**

1.指针变量一定接收的是地址值

2.指针变量的地址必须匹配对应的变量类型，int变量只能用 *int 来接收

3.我们先看个例子，大家可以先猜测一下最后输出的这个age是多少

```go
//demo09
func main() { 
	//定义一个int类型的变量
	var age int = 18
	a(age)
	fmt.Printf("函数修改后age的值：%v \n",age)
}

//将age的值修改为20
func a(age int)  {
	age = 20
	fmt.Println("a函数内的age值：",age)
}
```

**为了让大家加深印象，分享时再给出结论**

#####  数组

定义数组的语法：`var array[size] int `，数组中默认填充的都是对应数据类型的默认值

```go
//定义一个空数组,默认填充对应类型的默认值
var array01 [10] int
//给数组第一个元素赋值
array01[0] = 2
fmt.Println("array01:",array01)

//定义一个长度为5的数组，并给定值
array02 := [5] int {1,2,3,4,5}
fmt.Println("array02:",array02)

//定义一个未知长度的数组，编译器会根据元素推断出数组长度
//编译完成后，假设现在是10个元素，后期在方法中也不能给该数组加元素
//你可以理解为编译时，go帮你把...替换成了10就行
array03 := [...] int {1,3,5,6,9,7,5,2,2,6}
fmt.Printf("array中元素个数 = %v \n",len(array03))
```

```shell
array01: [2 0 0 0 0 0 0 0 0 0]
array02: [1 2 3 4 5]
array中元素个数 = 10 
```

#####  切片(Slice)

切片是对数组的抽象，由于数组定义好了长度，有局限性，所以就设计了切片，底层还是数组，切片会自动扩容，大家可以理解为java中的List

**切片的定义**

```go
//方式一:定义一个空数组 //demo11
var slice1 [] int
fmt.Println("slice1",slice1)
//方式二：通过make函数来 创建 make(type,len,cap) 切片类型，切片初始长度，切片初始容量
//默认填充对应类型的默认值
slice2 := make([]int,0,3)
fmt.Println("slice2",slice2)
//区别 slice1=nil slice2相当于初始化了，不为nil

//定义带初始值的切片，支持自动扩容
slice3 := [] int {1,2,3,4}
fmt.Println("slice3：",slice3)
```

```shell
slice []
slice2 []
slice3： [1 2 3 4]
```

**添加删除切片中的值**

go语言提供了切片添加元素的函数`append`，但是并没有提供删除元素的函数，不过我们利用现有的语法能自己删除删除的功能

```go
//给slice3添加元素  //demo11
slice3 = append(slice3, 1)
fmt.Println("slice3：",slice3)
//删除slice3的第一个元素
slice3 = remove(slice3, 0)
fmt.Println("slice3：",slice3)

/*
移除下标为index的元素，下面有将append的切片带的参数是啥意思
*/
func remove(slice [] int,index int) [] int {
	return append(slice[:index],slice[index + 1:]...)
}
```

```shell
slice3： [1 2 3 4 1]
slice3： [2 3 4 1]
```

**获取切片长度和容量**

使用len()方法和cap方法可以获取切片长度和容量

```go
//获取切片的长度
lenth := len(slice3)
//获取切片的容量
capsize := cap(slice3)
fmt.Println("切片的长度：",lenth," 容量：",capsize)
```

```shell
切片的长度： 4  容量： 8
```

**数组转切片**

数组转切片,传的是数组的引用，数组值变化，切片中对应的值也会变化

语法：`slice := array[startIndex,endIndex]` **左闭右开** ，也可以只传一个参数

`slice := array[:]`取数组的全部元素

`slice := array[startIndex:]` 从startIndex开始取，后面所有的元素

`slice := array[:endIndex]` 从第一个开始取，取到下标(endIndex-1)为止

```go
//定义一个数组  //demo12
array := [10] int {1,2,3,4,5,6,7,8,9,10}
//数组转对象，取数组中所有元素
slice1 := array[:]
fmt.Println("slice1:",slice1)
//数组转对象，从第一个开始取
slice2 := array[1:]
fmt.Println("slice2:",slice2)
//数组转对象，从第一个开始取到 5 - 1 个
slice3 := array[:5]
fmt.Println("slice3:",slice3)
```

```shell
slice1: [1 2 3 4 5 6 7 8 9 10]
slice2: [2 3 4 5 6 7 8 9 10]
slice3: [1 2 3 4 5]
```

**数组转切片传的是引用，切片的值会跟着数组的变化而变化**,如果后面切片添加元素超过了数组的长度时，底层会该切片创建一个容量为原来两倍的数组，此时老数组的值变动不会影响到扩容后的切片

```go
//修改数组最后一个元素为100
array[9] = 100
fmt.Println("slice1:",slice1)
//为slice切片添加一个元素，此时会扩容，底层新创建一个容量为2倍的数组
slice1 = append(slice1, 11)
//改变老数组的值
array[8] = 120
fmt.Println("slice1:",slice1)
```

```shell
slice1: [1 2 3 4 5 6 7 8 9 100]
slice1: [1 2 3 4 5 6 7 8 9 100 11]
```

那么除了数组扩容，还有别的办法让切片中的元素不随数组变动而变动呢？

```go
//创建一个切片 //demo12
slice4 := make([]int, 10)
//将切片slice1元素拷贝给slice4
copy(slice4,slice1)
//修改数组中的值
array[8] = 110
fmt.Println("slice1:",slice1)
fmt.Println("slice4:",slice4)
```

```shell
slice1: [1 2 3 4 5 6 7 8 110 100]
slice4: [1 2 3 4 5 6 7 8 9 100]
```

通过`copy`函数，我们可以发现数组变动以后，切片slice1的值变动了，但是拷贝的slice4还是拷贝之前的值，

我们可以发现copy是值拷贝，并不是引用拷贝。

#####  Map

Map是一种无序的键值对的集合。底层也是基于hash表来实现的，我们可以理解为Java中的HashMap

**Map定义**

```go
//定义一个map变量 初始值为nil，没有初始化不能往集合中添加元素
var map1 map[string] int
fmt.Println("map1",map1)
//异常：panic: assignment to entry in nil map
//map1["key"] = 1
//定义一个带初始值的map
map2 := map[string] int {"a": 1, "b":2}
fmt.Println("map2",map2)

//定义一个初始化的空map，容量为10
map3 := make(map[string]int, 10)
fmt.Println("map3",map3)
```

```shell
map1 map[]
map2 map[a:1 b:2]
map3 map[]
```

**元素添加**

`map[key] = value`，跟数据赋值差不多，数组给的是索引，map给的是key而已

```go
//map2添加元素 key = a  value = 1
map3["c"] = 1
map3["d"] = 2
fmt.Println("map3",map3)
```

```shell
map3 map[c:1 d:2]
```

**元素获取**

`value,ok := map[key]`，第一个元素为key对应的值，第二个表示`key`是否存在

```go
//通过key获取value 会返回一个元组
//第一个元素为key对应的值，第二个表示是否存在
value,ok := map3["d"]
fmt.Println("value:",value,"ok",ok)
```

```shell
value: 2 ok true
```

**元素删除**

通过`delete`函数，通过key删除

```shell
//删除map3中的元素
delete(map3,"d")
fmt.Println("map3",map3)
```

```shell
map3 map[c:1]
```

#####  管道

#####  函数

#####  接口

### 流程控制

#### 分支结构

##### 单分支

1.小括号可以省略，变量和if关键词之间一定要有空格

2.`{}`必须写，就算只有一行代码也不能省略

```go
flag := true //demo 14
//小括号可以省略，变量和if关键词之间一定要有空格
//{}必须写，就算只有一行代码也不能省略
//单分支
if flag {
    fmt.Println("flag：",flag)
}
```

##### 双分支

只需要注意单分支的两点，其余的跟Java用法一样

```go
//demo 14
//加入种子资源库，不然每次生成的随机数都是一样的
//因为rand是个伪随机生成器，而同一个共享源生成的随机数序列是恒定的
rand.Seed(time.Now().UnixNano())
num := rand.Intn(20)
//双分支
if num < 10 {
    fmt.Printf("随机数【%v】小于10",num)
} else {
    fmt.Printf("随机数【%v】大于10",num)
}
```

##### 多分支

```java
//demo 14
//多分支
if num < 5 {
    fmt.Printf("随机数【%v】小于5",num)
} else if num < 10 {
    fmt.Printf("随机数【%v】大于5小于10",num)
} else {
    fmt.Printf("随机数【%v】大于10",num)
}
```

##### switch

```go
//demo15
//随机分数除以10，得到评分结果10、9 评分为A  8、7 评分为B  6、5评分为C 4、3评分为D 小于3的劝退吧
rand.Seed(time.Now().UnixNano())
//随机一个分数
score := rand.Intn(101)
fmt.Println("你的分数为：",score)
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
```

1.switch后是一个表达式(即:常量值、变量、一个有返回值的函数等都可以

2.case后面的值如果是常量值(字面量)，则要求不能重复

3.case后的各个值的数据类型，必须和 switch 的表达式数据类型一致

4.case后面可以带多个值，使用逗号间隔。比如 case 值1,值2...

5.case后面不需要带break

6.default语句不是必须的，位置也是随意的。

7.switch后也可以不带表达式，当做if分支来使用

8.switch后也可以直接声明/定义一个变量，分号结束，不推荐

9.switch穿透，利用fallthrough关键字，如果在case语句块后增加fallthrough ,则会继续执行下一个case,也叫switch穿透。

#### 循环结构

```go
//类似于java里面的while
i := 0 //变量的初始化
for i < 5 { //条件表达式,判断条件
    fmt.Println(i)
    i++//累加
}
//死循环
for {
    
}
```

##### for range

 ```go
//demo16
//for range 有点像java里面的foreach
str := "hello"
//遍历的是索引和值
for index, value := range str {
    fmt.Printf("索引:%v, value:%c \n",index,value)
}
 ```

```shell
索引:0, value:h 
索引:1, value:e 
索引:2, value:l 
索引:3, value:l 
索引:4, value:o
```

##### 关键字

**break**

用法和java里面一样，不过go语言里面能通过加标签中断外层的循环

```go
//给外层循环设置一个label //demo16
label:
for i := 0; i < 5; i++ {
    for j := 2; j < 4; j++ {
        fmt.Printf("i:%v, j:%v",i,j)
        if i == 2 && j == 2 {
            break label //结束指定标签对应的循环，如果标签定义了，不用会报错
        }
    }
}
```

```shell
i:0, j:2 
i:0, j:3 
i:1, j:2 
i:1, j:3 
i:2, j:2 
```

**continue**

用法和java一样，不过跟broker一样，也能加标签，让对应标签的循环直接开始下一次循环

```go
label1:  //demo16
for i := 0; i < 3; i++ {
    for j := 2; j < 4; j++ {
        fmt.Printf("i:%v, j:%v \n",i,j)
        if i == 2 && j == 2 {
            continue label1 //结束指定标签对应的循环，如果标签定义了，不用会报错
        }
    }
}
fmt.Println("------ok------")
```

```shell
i:0, j:2 
i:0, j:3 
i:1, j:2 
i:1, j:3 
i:2, j:2 
i:0, j:2 
i:0, j:3 
i:1, j:2 
i:1, j:3 
i:2, j:2 
------ok------
```

**goto**

可以无条件的跳到某个标签的未知开始执行代码，通过下面的代码我们能知道，goto关键词下面那句代码直接呗跳过了

```go
//demo16
fmt.Println("开始执行goto关键词")
goto label2
fmt.Println("这里的内容将直接被跳过")
label2:
   fmt.Println("goto跳转过来后执行的代码")
```

```shell
开始执行goto关键词
goto跳转过来后执行的代码
```

**return**

等同于java里面的return，结束当前函数

### 函数

#### 函数定义

**语法**

func   函数名（形参列表)（返回值类型列表）{
执行语句..
return + 返回值列表
}

```go
//demo16
/*
  两数相加
  无返回值函数
 */
func add1(num1 int,num2 int)  {
	fmt.Printf("%v + %v = %v\n",num1,num2,num1 + num2)
}

/*
 数相减
 返回一个值
 */
func reduce(num1 int,num2 int) int {
	return num1 - num2
}

/**
  同时做加法和减法
  有返回多个值函数
 */
func addAndReduce(num1 int,num2 int) (int,int)  {
	sum := num1 + num2
	result := num1 - num2
	return sum,result
}

/**
  同时做加法和减法，返回值取别名
  有返回多个值函数
*/
func addAndReduce2(num1 int,num2 int) (sum int, result int)  {
	sum = num1 + num2
	result = num1 - num2
    //return后面就无须跟返回值变量名了
	return
}
/*
程序入口
*/
func main() {
	//调用无返回值
	add1(1,5)
	//一个返回值对象
	rResult := reduce(10, 3)
	fmt.Println("10 - 5 = ",rResult)
	//多个返回值对象
	addResult,reduceResult := addAndReduce(20,10)
	fmt.Println("20 + 10 = ",addResult,"\n20 - 10 = ",reduceResult)
	//调用取了别名的函数  只关心加法，用_忽略减法
	result, _ := addAndReduce2(30, 20)
	fmt.Println("30 + 20 = ",result)
}
```

```shell
1 + 5 = 6
10 - 5 =  7
20 + 10 =  30 
20 - 10 =  10
30 + 20 =  50
```

**友情提醒**

1.go函数不支持重载

2.函数名大写，相当于java里面的public 外部包也能调用到，小写相当于private

3.函数支持同时返回多个值，返回值可以取别名

4.多个返回值的情况可以用_忽略某个结果

5.基本数据类型都是值传递，在函数内修改，也不会影响到变量原本的值，这一点跟java里面不一样，但是可以通过传指针达到传引用的效果

6.函数本身也是一种数据类型，也可以赋值给变量，也可以通过变量来调用该函数

```go
//demo18
/*
  修改变量的值
 */
func test(age int)  {
	age = 20
	fmt.Println("test函数age = ",age)
}

/*
  修改变量的值，通过指针改变变量的值
*/
func test2(age *int)  {
	*age = 25
	fmt.Println("test2函数age = ",age)
}

/*
接收函数参数的函数
 */
func test3( age int,test func(int))  {
	test(age)
}

/*
程序入口
 */
func main() {
	//函数赋值给变量
	a:= test
	//定义年龄变量
	age := 18
	//变量也可以当函数调用
	a(age)
	fmt.Println("main函数中的age = ",age)
	//通过指针改变age的值
	test2(&age)
	fmt.Println("main函数中的age = ",age)
	//调用函数有形参为函数的函数
	test3(18,test)
}
```




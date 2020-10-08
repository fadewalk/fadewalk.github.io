---
title: FirstPage
date: 2020-10-07 15:19:24
tags:
---

## 编写第一个 Go 程序

### 开发环境构建

GOPATH

1. 在 1.8 版本前必须设置这个环境变量量
2. 1.8 版本后（含 1.8）如果没有设置使⽤用默认值
在 Unix 上默认为 $HOME/go , 在 Windows 上默认为 %USERPROFILE%/go
在 Mac 上 GOPATH 可以通过修改 ～/.bash_profile 来设置

```go
package main //包，表明代码所在的模块（包）
import "fmt" //引⼊入代码依赖
//功能实现
func main() {
fmt.Println("Hello World!")
}
```

### 应用程序入口
```
1.必须是main包：package main
2.必须是main方法：func main()
3.文件名不一定是 main.go
```
### 退出返回值

与其他主要编程语言的差异
·Go中main函数不支持任何返回值
·通过os.Exit来返回状态

### 获取命令行参数

与其他主要编程语言的差异
·main函数不支持传入参数

func main ~~(arg []string)~~

在程序中直接通过 os.Args获取命令行参数



### 第一个程序

命令行读取参数

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		fmt.Println("Hello World", os.Args[1])
	}
}
```



### 关键字

|     | 关键字      | 作用     | 一级分类 | 二级分类       | 三级分类   |
| --- | ----------- | -------- | -------- | -------------- | ---------- |
| 1   | var         | 变量声明 | 基本结构 | 变量与常量     | -          |
| 2   | const       | 常量声明 | 基本结构 | 变量与常量     | -          |
| 3   | package     | 包声明   | 基本结构 | 包管理         | -          |
| 4   | import      | 包引用   | 基本结构 | 包管理         | -          |
| 5   | func        | 函数声明 | 基本组件 | 函数           | -          |
| 6   | return      | 函数返回 | 基本组件 | 函数           | -          |
| 7   | interface   | 接口     | 基本组件 | 自定义类型     | -          |
| 8   | struct      | 结构体   | 基本组件 | 自定义类型     | -          |
| 9   | type        | 定义类型 | 基本组件 | 自定义类型     | -          |
| 10  | map         |          | 基本组件 | 引用类型       | -          |
| 11  | range       |          | 基本组件 | 引用类型       | -          |
| 12  | go          |          | 流程控制 | 并发           | -          |
| 13  | select      |          | 流程控制 | 并发           | -          |
| 14  | chan        |          | 流程控制 | 并发           | -          |
| 15  | if          |          | 流程控制 | 单任务流程控制 | 单分支流程 |
| 16  | else        |          | 流程控制 | 单任务流程控制 | 单分支流程 |
| 17  | switch      |          | 流程控制 | 单任务流程控制 | 多分支流程 |
| 18  | case        |          | 流程控制 | 单任务流程控制 | 多分支流程 |
| 19  | default     |          | 流程控制 | 单任务流程控制 | 多分支流程 |
| 20  | fallthrough |          | 流程控制 | 单任务流程控制 | 多分支流程 |
| 21  | for         |          | 流程控制 | 单任务流程控制 | 循环流程   |
| 22  | break       |          | 流程控制 | 单任务流程控制 | 循环流程   |
| 23  | continue    |          | 流程控制 | 单任务流程控制 | 循环流程   |
| 24  | goto        |          | 流程控制 | 单任务流程控制 | -          |
| 25  | defer       |          | 流程控制 | 延时流程控制   | -          |



## 变量与常量

>The master has failed more times than the beginner has tried.

### 编写测试程序

1. 源码文件以 _test 结尾：xxx_test.go
2. 测试方法名以 Test 开头：func TestXXX(t *testing.T) {…}

### 变量量赋值

与其他主要编程语言的差异
·赋值可以进行自动类型推断
·在一个赋值语句中可以对多个变量进行同时赋值

斐波拉契数列

```go
package ch2

import "testing"

func TestFibList(t *testing.T) {
   /* 变量命名的方式1
      var a int = 1
      var b int = 1*/

   /* 方式2
      var (
         a int = 1
         b     = 2
      )*/

   //方式3
   a := 1
   b := 1
   t.Log("", a)
   for i := 0; i < 5; i++ {
      t.Log("", b)

      /*    tmp := a
            a = b
            b = tmp + a*/

      a, b = b, a+b //交换赋值
   }
}

func TestExchange(t *testing.T) {
   a := 1
   b := 2
   /*tmp := a
   a = b
   b = tmp
   */
   a, b = b, a //赋值交换
   t.Log(a, b)
}
```

### 常量定义

与其他主要编程语言的差异
快速设置连续值

```go
package ch2

import "testing"

const (
	Monday = 1 + iota
	Tuesday
	Wednesday
	//批量递增命名
)

func TestConstantTry(t *testing.T) {
	t.Log(Monday, Tuesday, Wednesday)
}

/*
=== RUN   TestConstantTry
constant_test.go:13: 1 2 3
--- PASS: TestConstantTry (0.00s)
PASS*/

const (
	Readable = 1 << iota
	Writable
	Executable
)

func TestConstant2(t *testing.T) {
	t.Log(Readable, Writable, Executable)
	a := 1 //0001
	t.Log(a&Readable == Readable, a&Writable == Writable, a&Executable == Executable)

	//	位运算
}

/*=== RUN   TestConstant2
constant_test.go:29: 1 2 4
constant_test.go:31: true false false
--- PASS: TestConstant2 (0.00s)
PASS*/
```



## 数据类型

### 基本数据类型

```
bool
string
int int8 int16 int32 int64
uint uint8 uint16 uint32 uint64 uintptr
byte // alias for uint8
rune // alias for int32,represents a Unicode code point
float32 float64
complex64 complex128
```

![image-20201006104740280](https://gitee.com/wyule/picture-bed/raw/master/assets/image-20201006104740280.png)

### 类型转化

与其他主要编程语言的差异

1.Go语言不允许隐式类型转换

2.别名和原有类型也不能进行隐式类型转换

### 类型的预定义值

1. math.MaxInt64
2. math.MaxFloat64
3. math.MaxUint32

### 指针类型

与其他主要编程语言的差异

1. 不支持指针运算
2. string是值类型，其默认的初始化值为空字符串，而不是nil

```go
package type_test

import "testing"

type MyInt int64

func TestImplict(t *testing.T) {
   var a int32 = 1
   var b int64
   //b = a // 不能隐式转换cannot use a (type int32) as type int64 in assignment
   b = int64(a)
   t.Log(b)

   var c MyInt
   c = MyInt(b) // 别名类型也要显示转换

   t.Log(a, b, c)

}

func TestPoint(t *testing.T) {
   a := 1
   aPtr := &a // 获取a的指针
   t.Log(a, aPtr)
   //aPtr = aPtr + 1        // 指针不能进行运算
   t.Logf("%T %T", a, aPtr)

}

/*=== RUN   TestPoint
type_test.go:24: 1 0xc000094268
type_test.go:26: int *int
--- PASS: TestPoint (0.00s)
PASS*/
func TestString(t *testing.T) {
   var s string
   t.Log("*" + s + "*") // //初始化零值是"" 是空字符串而不是nil
   t.Log(len(s))
}

/*
=== RUN   TestString
type_test.go:37: **
type_test.go:38: 0
--- PASS: TestString (0.00s)
PASS
*/
```



## 运算符



### 算术运算符

| 运算符 | 描述 | 实例               |
| ------ | ---- | ------------------ |
| +      | 相加 | A + B 输出结果 30  |
| -      | 相减 | A - B 输出结果 -10 |
| *      | 相乘 | A * B 输出结果 200 |
| /      | 相除 | B / A 输出结果 2   |
| %      | 求余 | B % A 输出结果 0   |
| ++     | 自增 | A++ 输出结果 11    |
| --     | 自减 | A-- 输出结果 9     |

Go 语⾔言没有前置的 ++，- -，~~（++a）~~

### 比较运算符

![image-20201006182638378](https://gitee.com/wyule/picture-bed/raw/master/assets/image-20201006182638378.png)



### 用==比较数组

相同维数且含有相同个数元素的数组才可以比较

每个元素都相同的才相等

### 逻辑运算符

![image-20201006182806017](https://gitee.com/wyule/picture-bed/raw/master/assets/image-20201006182806017.png)

### 位运算符

![image-20201006183106374](https://gitee.com/wyule/picture-bed/raw/master/assets/image-20201006183106374.png)





与其他主要编程语⾔言的差异
&^ 按位置零
```
1 &^ 0 -- 1
1 &^ 1 -- 0
0 &^ 1 -- 0
0 &^ 0 -- 0
```


只要右边位数设为1，左边清零都为0
右边位数设为0，左边的原来是什么，结果就是什么

```go
package ch4

import "testing"

const (
   Readable   = 1 << iota
   Writable   // 2
   Executable //4
)

func TestCompareArray(t *testing.T) {
   a := [...]int{1, 2, 3, 4}
   b := [...]int{1, 3, 2, 5}
   //c := [...]int{1, 2, 3, 4, 5}
   d := [...]int{1, 2, 3, 4}

   t.Log(a == b) //false
   //t.Log(a == c)  // 长度不一样不能进行对比
   t.Log(a == d) //true
   t.Log(Readable, Writable, Executable)
}

func TestBitClear(t *testing.T) {
   a := 7 //0111
   a = a &^ Readable
   a = a &^ Executable
   t.Log(a&Readable == Readable, a&Writable == Writable, a&Executable == Executable)
   // false true false
}
```



## 编写结构化程序

### 循环

与其他主要编程语言的差异
Go语言仅支持循环关键字for

for ~~(~~ j := 7; j <= 9; j++ ~~)~~

![image-20201007161817470](https://gitee.com/wyule/picture-bed/raw/master/assets/image-20201007161817470.png)

### if 条件

```
if  condition {
    // code to be executed if condition is true
} else {
    // code to be executed if condition is false
}
if  condition-1 {
    // code to be executed if condition-1 is true
} else if condition-2 {
    // code to be executed if condition-2 is true
} else {
    // code to be executed if both condition1 and condition2 are false
}
```

与其他主要编程语言的差异
1. **condition表达式结果必须为布尔值**
2. 支持变量賦值：

```
if  var declaration;  condition {
    // code to be executed if condition is true
}
```

### switch 条件

![image-20201007162138298](https://gitee.com/wyule/picture-bed/raw/master/assets/image-20201007162138298.png)



与其他主要编程语言的差异

1. 条件表达式不限制为常量或者整数；
2. 单个 case 中，可以出现多个结果选项, 使用逗号分隔；
3. 与 C 语⾔言等规则相反，Go 语言不需要用break来明确退出一个 case；
4. 可以不设定 switch 之后的条件表达式，在此种情况下，整个 switch 结构与多个 if…else… 的逻辑作用等同

```go
package ch5

import "testing"

func TestWhileLoop(t *testing.T) {

   for n := 0; n < 5; n++ {
      t.Log(n)
      //n++
   }
}

/*=== RUN   TestWhileLoop
condition_test.go:8: 0
condition_test.go:8: 1
condition_test.go:8: 2
condition_test.go:8: 3
condition_test.go:8: 4
--- PASS: TestWhileLoop (0.00s)
PASS*/

func TestIf(t *testing.T) {
   if a := 1 == 1; a {
      t.Log("1==1")
   }

   /* if v, err := someFun(); err = nil{
         t.Log("1==1"),
      } else {
         t.Log("1==1")
      }*/

}

func TestSitchCase(t *testing.T) {
   for i := 0; i < 5; i++ {
      switch i {
      case 0, 2:
         t.Log("Even")
      case 1, 3:
         t.Log("Odd")
      default:
         t.Log("it is not 0-3")
      }
   }
}

func TestSitchConditionCase(t *testing.T) {
   for i := 0; i < 5; i++ {
      switch {
      case i%2 == 0:
         t.Log("Even")
      case i%2 == 1:
         t.Log("Odd")
      default:
         t.Log("unkown")
      }
   }
}
```

## 数组和切片

### 数组的声明

```
var a [3]int //声明并初始化为默认零值
a[0] = 1
b := [3]int{1, 2, 3} //声明同时初始化
c := [2][2]int{{1, 2}, {3, 4}} //多维数组初始化  
```

### 数组元素遍历

与其他主要编程语言的差异

```
func TestTravelArray(t *testing.T) {
	a := [...]int{1, 2, 3, 4, 5} //不指定元素个数
	for idx/*索引*/, elem/*元素*/ := range a {
		fmt.Println(idx, elem)
}
}
```

### 数组截取

a[开始索引(包含), 结束索引(不包含)]  
索引不能是负数

```
a := [...]int{1, 2, 3, 4, 5}
a[1:2] //2
a[1:3] //2,3
a[1:len(a)] //2,3,4,5
a[1:] //2,3,4,5
a[:3] //1,2,3
```

### 切片内部结构

![image-20201007212611819](https://gitee.com/wyule/picture-bed/raw/master/assets/image-20201007212611819.png)

### 切片声明

```
var s0 []int
s0 = append(s0, 1)

s := []int{}
s1 := []int{1, 2, 3}
s2 := make([]int, 2, 4)
/*[]type, len, cap
其中len个元素会被初始化为默认零值，未初始化元素不不可以访问
*/
```



### 切片共享存储结构

python 列表切片是共享的数据，

<img src="https://gitee.com/wyule/picture-bed/raw/master/assets/image-20201007213737043.png" alt="image-20201007213737043" style="zoom: 33%;" />

### 数组 vs 切片

1.容量是否可伸缩
2.是否可以进行比较

### paly code

```go
package ch6

import "testing"

func TestSliceInt(t *testing.T) {
	var s0 []int
	t.Log(len(s0), cap(s0))

	s0 = append(s0, 1) //  slice_test.go:9: 1 1
	t.Log(len(s0), cap(s0))

	s1 := []int{1, 2, 3, 4}
	t.Log(len(s1), cap(s1))

	s2 := make([]int, 3, 5)
	t.Log(len(s2), cap(s2))
	//t.Log(s2[0], s2[1], s2[2],s2[3], s2[4]) // 异常
	t.Log(s2[0], s2[1], s2[2]) //     slice_test.go:18: 0 0 0
	s2 = append(s2, 1)
	t.Log(s2[0], s2[1], s2[2], s2[3])
	t.Log(len(s2), cap(s2))
}

/*=== RUN   TestSliceInt
    slice_test.go:7: 0 0
    slice_test.go:10: 1 1
    slice_test.go:13: 4 4
    slice_test.go:16: 3 5
    slice_test.go:18: 0 0 0
    slice_test.go:20: 0 0 0 1
    slice_test.go:21: 4 5
--- PASS: TestSliceInt (0.00s)
PASS
*/

func TestSliceGrowing(t *testing.T) {
	s := []int{}
	for i := 0; i < 10; i++ {
		s = append(s, i)
		t.Log(len(s), cap(s))
	}
} /*
=== RUN   TestSliceGrowing
slice_test.go:40: 1 1
slice_test.go:40: 2 2
slice_test.go:40: 3 4
slice_test.go:40: 4 4
slice_test.go:40: 5 8
slice_test.go:40: 6 8
slice_test.go:40: 7 8
slice_test.go:40: 8 8
slice_test.go:40: 9 16
slice_test.go:40: 10 16			//最初没有开始定义cap，因此cap是自增的，以满足slice的append操作
--- PASS: TestSliceGrowing (0.00s)
PASS*/

func TestSliceShareMemory(t *testing.T) {
	year := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep",
		"Oct", "Nov", "Dec"}
	Q2 := year[3:6]
	t.Log(Q2, len(Q2), cap(Q2))
	summer := year[5:8]
	t.Log(summer, len(summer), cap(summer))
	summer[0] = "Unkown"
	t.Log(Q2)
	t.Log(year)

}

/*
=== RUN   TestSliceShareMemory
slice_test.go:48: [Apr May Jun] 3 9
slice_test.go:50: [Jun Jul Aug] 3 7
slice_test.go:52: [Apr May Unkown]
slice_test.go:53: [Jan Feb Mar Apr May Unkown Jul Aug Sep Oct Nov Dec]
--- PASS: TestSliceShareMemory (0.00s)
PASS*/

func TestSliceComparing(t *testing.T) {
	//a := []int{1, 2, 3, 4}
	//b := []int{1, 2, 3, 4}
	//if a == b { //invalid operation: a == b (slice can only be compared to nil)
	//	t.Log("equal")
	//}
}
```



## Map基础

### Map 声明

```
m := map[string]int{"one": 1, "two": 2, "three": 3}
m1 := map[string]int{}
m1["one"] = 1
m2 := make(map[string]int, 10 /*Initial Capacity*/)
//为什什么不不初始化len？  
```

### Map 元素的访问

在访问的 Key 不不存在时，仍会返回零值，不不能通过返回 nil 来判断元素是否存在

```
if v, ok := m["four"]; ok {
	t.Log("four", v)
	} else {
	t.Log("Not existing")
}
```

与其他主要编程语⾔言的差异  

### Map 遍历

```
m := map[string]int{"one": 1, "two": 2, "three": 3}
for k, v := range m {
	t.Log(k, v)
}
```

​	

```go
package ch7

import (
	"testing"
)

func TestIntMap(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 2, 3: 9}
	t.Log(m1[2])
	t.Logf("len m1=%d", len(m1))

	m2 := map[int]string{}
	m2[4] = "lewen"
	t.Logf("len m2=%d", len(m2))

	m3 := make(map[int]int, 10)
	t.Logf("len m3=%d", len(m3)) // ,cap(m3)) 不可以
	/*
		=== RUN   TestIntMap
		map_test.go:7: 2
		map_test.go:8: len m1=3
		map_test.go:12: len m2=1
		map_test.go:15: len m3=0
		--- PASS: TestIntMap (0.00s)
		PASS
	*/
}

func TestAccessNotExistingKey(t *testing.T) {
	m1 := map[int]int{}
	t.Log(m1[1]) //不存在反回0值
	m1[2] = 0
	t.Log(m1[2])
	if v, bol := m1[3]; bol {
		t.Log(bol)
		t.Logf("key 3's value is %d", v)
	} else {
		t.Log(bol)
		t.Log("key 3 is not existing.")
	}
	/*	map_test.go:31: 0
		map_test.go:33: 0
		map_test.go:38: false
		map_test.go:39: key 3 is not existing.*/
}

func TestTravelMap(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 2, 3: 9}
	for k, v := range m1 {
		t.Log(k, v)
	}
	/*map_test.go:47: 1 1
	  map_test.go:47: 2 2
	  map_test.go:47: 3 9*/
}
```

### Map 扩展

#### Map与工厂模式

·Map的value可以是一个方法
·与Go的 Dock type接口方式一起，可以方便的实现单一方法对象的工厂模式

```go
func TestMapWithFunValue(t *testing.T) {
	m := map[int]func(arg int) int{}
	m[1] = func(arg int) int { return arg }
	m[2] = func(arg int) int { return arg * arg }
	m[3] = func(arg int) int { return arg * arg * arg }
	t.Log(m[1](2), m[2](2), m[3](2))
	//   map_test.go:56: 2 4 8
}
```

#### 实现Set

Go的内置集合中没有Set实现，可以 map[type]bool
1.元素的唯一性

2.基本操作
1）添加元素
2）判断元素是否存在
3）删除元素
4）元素个数

```go

func TestMapForSet(t *testing.T) {
	mySet := map[int]bool{}
	mySet[1] = true
	n := 3

	if mySet[n] {
		t.Logf("%d is existing", n)
	} else {
		t.Logf("%d is not existing ", n)
	}
	mySet[3] = true
	t.Log(len(mySet))

	delete(mySet, 1)
	/*	map_test.go:76: 3 is not existing
		map_test.go:79: 2*/
}

```

## 字符串与字符编码








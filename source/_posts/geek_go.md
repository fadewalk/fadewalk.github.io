---
title: 我的Go学习笔记
date: 2020-10-07 15:19:24
tags: go
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

bool
string
int int8 int16 int32 int64
uint uint8 uint16 uint32 uint64 uintptr
byte // alias for uint8
rune // alias for int32,represents a Unicode code point
float32 float64
complex64 complex128

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










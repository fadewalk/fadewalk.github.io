package type_test

import (
	"math"
	"testing"
)

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

	num := math.MaxFloat64
	t.Log(num)
}

/*
=== RUN   TestString
type_test.go:37: **
type_test.go:38: 0
--- PASS: TestString (0.00s)
PASS
*/

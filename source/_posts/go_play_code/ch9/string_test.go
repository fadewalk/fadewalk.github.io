package ch9

import (
	"testing"
	"unsafe"
)

func TestString(t *testing.T) {
	var s string
	t.Log(s)
	s = "hello"
	t.Log(len(s))

	//s[1] ='2' //string是不可变的byte slice

	s = "\xE4\xB8\xA5" //可以存储任何二进制数据
	//s = "\xE4\xBA\xBB\xFF"
	t.Log(s)
	t.Log(len(s))

	s = "中"
	t.Log(len(s)) //是byte数

	c := []rune(s)
	t.Log(len(c))
	t.Log("rune size:", unsafe.Sizeof(c[0]))
	t.Logf("中 unicode %x", c[0])
	t.Logf("中 UTF8 %x", s)

}

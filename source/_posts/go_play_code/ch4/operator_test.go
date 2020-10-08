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

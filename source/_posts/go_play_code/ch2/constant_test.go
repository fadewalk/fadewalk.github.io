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

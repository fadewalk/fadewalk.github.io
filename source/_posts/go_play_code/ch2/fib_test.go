package ch2

import "testing"

func TestFibList(t *testing.T) {
	/*	变量命名的方式1
		var a int = 1
		var b int = 1*/

	/*	方式2
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

		/*		tmp := a
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

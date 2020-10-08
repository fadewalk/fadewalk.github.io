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

	/*	if v, err := someFun(); err = nil{
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

package ch6

import "testing"

func TestArrayInt(t *testing.T) {
	var arr [3]int
	t.Log(arr[1], arr[2])
	arr1 := [4]int{1, 2, 3, 4}
	arr2 := [...]int{1, 3, 4, 5}
	arr1[1] = 5
	t.Log(arr1[1], arr1, arr2)
}

/*
=== RUN   TestArrayInt
array_test.go:7: 0 0
array_test.go:11: 5 [1 5 3 4] [1 3 4 5]
--- PASS: TestArrayInt (0.00s)
PASS*/

func TestArrayTravel(t *testing.T) {
	arr3 := [...]int{1, 3, 4, 5}
	for i := 0; i < len(arr3); i++ {
		t.Log(arr3[i])
	}

	for _, e := range arr3 {
		t.Log(e)
	}

}

/*=== RUN   TestArrayTravel
array_test.go:24: 1
array_test.go:24: 3
array_test.go:24: 4
array_test.go:24: 5
array_test.go:28: 1
array_test.go:28: 3
array_test.go:28: 4
array_test.go:28: 5
--- PASS: TestArrayTravel (0.00s)
PASS*/

func TestArraySection(t *testing.T) {
	arr3 := [...]int{1, 3, 4, 5}
	arr3_sec := arr3[3:]
	//arr4_sec := arr3[-1:]  //不支持负数索引
	t.Log(arr3_sec)
}

/*=== RUN   TestArraySection
array_test.go:48: [5]
--- PASS: TestArraySection (0.00s)
PASS*/

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

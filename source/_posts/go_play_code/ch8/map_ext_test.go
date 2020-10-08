package ch8

import "testing"

func TestMapWithFunValue(t *testing.T) {
	m := map[int]func(arg int) int{}
	m[1] = func(arg int) int { return arg }
	m[2] = func(arg int) int { return arg * arg }
	m[3] = func(arg int) int { return arg * arg * arg }
	t.Log(m[1](2), m[2](2), m[3](2))
	//   map_test.go:56: 2 4 8
}

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

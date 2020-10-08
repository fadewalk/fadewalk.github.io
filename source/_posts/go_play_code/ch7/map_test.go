package ch7

import (
	"testing"
)

func TestIntMap(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 2, 3: 9}
	t.Log(m1[2])
	t.Logf("len m1=%d", len(m1))

	m2 := map[int]string{}
	m2[4] = "lewen"
	t.Logf("len m2=%d", len(m2))

	m3 := make(map[int]int, 10)
	t.Logf("len m3=%d", len(m3)) // ,cap(m3)) 不可以
	/*
		=== RUN   TestIntMap
		map_test.go:7: 2
		map_test.go:8: len m1=3
		map_test.go:12: len m2=1
		map_test.go:15: len m3=0
		--- PASS: TestIntMap (0.00s)
		PASS
	*/
}

func TestAccessNotExistingKey(t *testing.T) {
	m1 := map[int]int{}
	t.Log(m1[1]) //不存在反回0值
	m1[2] = 0
	t.Log(m1[2])
	if v, bol := m1[3]; bol {
		t.Log(bol)
		t.Logf("key 3's value is %d", v)
	} else {
		t.Log(bol)
		t.Log("key 3 is not existing.")
	}
	/*	map_test.go:31: 0
		map_test.go:33: 0
		map_test.go:38: false
		map_test.go:39: key 3 is not existing.*/
}

func TestTravelMap(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 2, 3: 9}
	for k, v := range m1 {
		t.Log(k, v)
	}
	/*map_test.go:47: 1 1
	  map_test.go:47: 2 2
	  map_test.go:47: 3 9*/
}

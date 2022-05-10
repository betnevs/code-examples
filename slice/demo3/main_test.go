package main

import (
	"sync"
	"testing"
)

func TestSlice(t *testing.T) {
	s := []int{}
	var wg sync.WaitGroup

	// 外部变量记录每个 goroutine append 的数量
	count := 0
	// 10 个 goroutine 并发 append 10000 个数字到 slice s 中，最终 s 正确长度为 10 * 10000 = 100000
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i, count int) {
			for j := 0; j < 10000; j++ {
				s = append(s, j)
				count++
			}
			t.Logf("G%d append count:%d\n", i, count)
			wg.Done()
		}(i, count)
	}
	wg.Wait()

	if len(s) != 100000 {
		t.Errorf("s.len:%d != 100000", len(s))
	}
}

func TestMap(t *testing.T) {
	m := map[int]int{}
	var wg sync.WaitGroup

	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func(i int) {
			if _, exists := m[i]; !exists {
				m[i] = i
			}
			wg.Done()
		}(i)
	}
	wg.Wait()

	t.Log(m)
}

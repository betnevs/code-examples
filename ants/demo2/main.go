package main

import (
	"fmt"
	"math/rand"
	"sync"

	"github.com/panjf2000/ants/v2"
)

const (
	DataSize    = 10000
	DataPerTask = 100
)

type Task struct {
	index int
	nums  []int
	sum   int
	wg    *sync.WaitGroup
}

func (t *Task) Do() {
	for _, num := range t.nums {
		t.sum += num
	}
	t.wg.Done()
}

type taskFunc func()

func taskFuncWrapper(nums []int, i int, sum *int, wg *sync.WaitGroup) taskFunc {
	return func() {
		for _, num := range nums[i*DataPerTask : (i+1)*DataPerTask] {
			*sum += num
		}
		fmt.Printf("task: %d sum: %d\n", i+1, *sum)
		wg.Done()
	}
}

func main() {
	p, _ := ants.NewPool(10)
	defer p.Release()
	nums := make([]int, DataSize, DataSize)
	for i := range nums {
		nums[i] = rand.Intn(1000)
	}
	var wg sync.WaitGroup
	wg.Add(DataSize / DataPerTask)
	partSums := make([]int, DataSize/DataPerTask, DataSize/DataPerTask)
	for i := 0; i < DataSize/DataPerTask; i++ {
		p.Submit(taskFuncWrapper(nums, i, &partSums[i], &wg))
	}
	wg.Wait()
	fmt.Printf("running goroutines: %d\n", ants.Running())
	var sum int
	for _, part := range partSums {
		sum += part
	}
	var expect int
	for _, num := range nums {
		expect += num
	}
	fmt.Printf("finish all goroutines, result is %d, expect is %d\n", sum, expect)
}

package main

import (
	"fmt"

	"github.com/panjf2000/ants/v2"
)

func tasks() []int {
	num := 1000
	res := make([]int, num)
	for i := 0; i < num; i++ {
		res[i] = i
	}
	return res
}

func main() {
	p, _ := ants.NewPool(10)
	defer p.Release()
	for _, task := range tasks() {
		task := task
		p.Submit(func() {
			fmt.Println(task)
		})
	}
	// 模拟服务
	select {}
}

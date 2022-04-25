package main

import (
	"fmt"
	"time"
)

type Filter func(c *Context)

type FilterBuilder func(next Filter) Filter

var _ FilterBuilder = MetricsFilterBuilder

func MetricsFilterBuilder(next Filter) Filter {
	return func(c *Context) {
		start := time.Now().Nanosecond()
		next(c)
		end := time.Now().Nanosecond()
		fmt.Println("cost: ", end-start)
	}
}

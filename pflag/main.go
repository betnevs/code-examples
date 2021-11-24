package main

import (
	"fmt"
	"time"

	"github.com/spf13/pflag"
)

func main() {
	var a int
	var b time.Duration
	var c bool
	pflag.IntVar(&a, "aa", 0, "usage")
	pflag.DurationVarP(&b, "bb", "b", 0, "time duration")
	pflag.BoolVar(&c, "cc", false, "get bool")
	pflag.Parse()
	//pflag.CommandLine.MarkDeprecated("aa", "delete aa")
	fmt.Println(a, b)
}

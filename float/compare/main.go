package main

import (
	"fmt"
	"math"
)

func main() {

	var f1 float32 = 16777216.0
	var f2 float32 = -16777217.0
	fmt.Println(f1 == f2) // true

	fmt.Printf("%d\n", math.Float32bits(f1))
	fmt.Printf("%d\n", math.Float32bits(f2))
}

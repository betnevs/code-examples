package main

import (
	"context"
	"fmt"
)

type ctxKey struct{}

type ctxKey2 struct{}

type ctxKey3 int

func main() {
	ck1 := ctxKey{}
	ck2 := ctxKey2{}
	parentCtx := context.Background()
	ctx := context.WithValue(parentCtx, ck1, "aaaa")
	fmt.Println(ctx.Value(ck1))

	fmt.Printf("%p, %p\n", &ck1, &ck2)
	fmt.Println(ck1 == ctxKey(ck2))

	ck3 := ctxKey3(11)

	ctx2 := context.WithValue(parentCtx, ck3, "bbb")
	fmt.Println(ctx2.Value(ck3), ctx2.Value(11))
}

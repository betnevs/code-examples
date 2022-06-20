package trie

import (
	"fmt"
	"strings"
	"testing"
)

func TestSplit(t *testing.T) {
	a := "aa/bb/cc"
	res := strings.Split(a, "/")
	fmt.Println(res, len(res))

}

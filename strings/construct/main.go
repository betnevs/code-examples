package main

import (
	"bytes"
	"fmt"
	"strings"
)

var s1 []string = []string{
	"Rob Pike ",
	"Robert Griesemer ",
	"Ken Thompson ",
}

func concatStringByOperator(s1 []string) string {
	var s string
	for _, v := range s1 {
		s += v
	}
	return s
}

func concatStringBySprintf(s1 []string) string {
	var s string
	for _, v := range s1 {
		s = fmt.Sprintf("%s%s", s, v)
	}
	return s
}

func concatStringByJoin(s1 []string) string {
	return strings.Join(s1, "")
}

func concatStringByStringsBuilder(s1 []string) string {
	var b strings.Builder
	for _, v := range s1 {
		b.WriteString(v)
	}
	return b.String()
}

func concatStringByStringsBuilderWithInitSize(s1 []string) string {
	var b strings.Builder
	b.Grow(64)
	for _, v := range s1 {
		b.WriteString(v)
	}
	return b.String()
}

func concatStringByBytesBuffer(s1 []string) string {
	var b bytes.Buffer
	for _, v := range s1 {
		b.WriteString(v)
	}
	return b.String()
}

func concatStringByBytesBufferWithInitSize(s1 []string) string {
	buf := make([]byte, 0, 64)
	b := bytes.NewBuffer(buf)
	for _, v := range s1 {
		b.WriteString(v)
	}
	return b.String()
}

func main() {

}

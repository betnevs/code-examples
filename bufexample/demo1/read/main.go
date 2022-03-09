package main

import (
	"bufio"
	"fmt"
	"io"
)

type Reader struct {
	counter int
}

func (r *Reader) Read(p []byte) (n int, err error) {
	fmt.Println("Read")
	if r.counter >= 3 {
		return 0, io.EOF
	}
	s := "a,b,c,d,e,f,g,h,i,j,k,l,m,n,o,p,q"
	copy(p, s)
	r.counter += 1
	return len(s), nil
}

func main() {
	r := new(Reader)
	br := bufio.NewReader(r)
	for {
		token, err := br.ReadString(',')
		fmt.Println("Token:%q\n", token)
		if err == io.EOF {
			fmt.Println("Read done")
			break
		}
	}
}

package main

type T struct {
	a int
}

func (t T) M1() {
	t.a = 10
}

func (t *T) M2() {
	t.a = 11
}

func main() {
	var t T
	println(t.a) // 0

	t.M1()
	println(t.a) // 0

	t.M2()
	println(t.a) // 11

	p := &t
	p.M2()
	println(t.a) // 11

}

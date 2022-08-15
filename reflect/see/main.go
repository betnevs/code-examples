package main

import (
	"bytes"
	"fmt"
	"reflect"
)

type User struct {
	Name    string
	Age     int
	Married bool
}

func inspectStruct(u interface{}) {
	v := reflect.ValueOf(u)
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		switch field.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			fmt.Printf("field: %d type: %s value: %d\n", i, field.Type().Name(), field.Int())
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			fmt.Printf("field: %d type: %s value: %d\n", i, field.Type().Name(), field.Uint())
		case reflect.Bool:
			fmt.Printf("field: %d type: %s value: %t\n", i, field.Type().Name(), field.Bool())
		case reflect.String:
			fmt.Printf("field: %d type: %s value: %q\n", i, field.Type().Name(), field.String())
		default:
			fmt.Printf("field: %d unhandled kind: %s\n", i, field.Kind())
		}
	}
}

func inspectMap(m interface{}) {
	v := reflect.ValueOf(m)
	for _, k := range v.MapKeys() {
		field := v.MapIndex(k)
		fmt.Printf("%v => %v\n", k.Interface(), field.Interface())
	}
}

func inspectSliceArray(sa interface{}) {
	v := reflect.ValueOf(sa)

	fmt.Printf("%c", '[')

	for i := 0; i < v.Len(); i++ {
		elem := v.Index(i)
		fmt.Printf("%v ", elem.Interface())
	}

	fmt.Printf("%c\n", ']')
}

func Add(a, b int) int {
	return a + b
}

func Greeting(name string) string {
	return "hello " + name
}

func invoke(f interface{}, args ...interface{}) {
	v := reflect.ValueOf(f)
	argV := make([]reflect.Value, 0, len(args))

	for _, arg := range args {
		argV = append(argV, reflect.ValueOf(arg))
	}

	rets := v.Call(argV)

	fmt.Println("ret:")

	for _, ret := range rets {
		fmt.Println(ret.Interface())
	}
}

func inspectFunc(name string, f interface{}) {
	t := reflect.TypeOf(f)
	fmt.Println("input: ", name)

	for i := 0; i < t.NumIn(); i++ {
		t := t.In(i)
		fmt.Print(t.Name())
		fmt.Print(" ")
	}

	fmt.Println()

	for i := 0; i < t.NumOut(); i++ {
		t := t.Out(i)
		fmt.Print(t.Name())
		fmt.Println(" ")
	}

	fmt.Println("\n===============")
}

func inspectMethod(o interface{}) {
	t := reflect.TypeOf(o)

	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		fmt.Println(m)
	}
}

func (u *User) SetName(n string) {
	u.Name = n
}

func (u *User) SetAge(a int) {
	u.Age = a
}

func main() {
	m := make(map[string]int, 10)
	fmt.Println(len(m))
	m["a"] = 1
	m["b"] = 2
	fmt.Println(len(m))

	a := [3]int{1, 2, 3}
	b := []int{9, 10}
	copy(b, a[:])
	fmt.Println(a, b)

	u := User{
		Name:    "DJ",
		Age:     18,
		Married: true,
	}
	inspectStruct(u)
	inspectStruct(bytes.Buffer{})

	fmt.Println("-------------------------")

	inspectMap(map[uint32]uint32{1: 2, 3: 4})

	fmt.Println("-------------------------")

	inspectSliceArray([]int{1, 2, 3})
	inspectSliceArray([3]int{11, 21, 31})

	fmt.Println("-------------------------")

	inspectFunc("Add", Add)
	inspectFunc("Greeting", Greeting)

	fmt.Println("-------------------------")

	u2 := User{
		Name: "yangjie",
		Age:  18,
	}

	inspectMethod(&u2)

	fmt.Println("-------------------------")
	invoke(Add, 1, 2)
	invoke(Greeting, "dj")

	fmt.Println("-------------------------")

	m1 := M{1, 2, '+'}
	m2 := M{3, 2, '-'}
	m3 := M{3, 2, '*'}
	m4 := M{3, 2, '/'}
	invoke(m1.Op)
	invoke(m2.Op)
	invoke(m3.Op)
	invoke(m4.Op)

}

type M struct {
	a, b int
	op   byte
}

func (m M) Op() int {
	switch m.op {
	case '+':
		return m.a + m.b
	case '-':
		return m.a - m.b
	case '*':
		return m.a * m.b
	case '/':
		return m.a / m.b
	default:
		panic("invalid op")
	}
}

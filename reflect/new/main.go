package main

import (
	"errors"
	"fmt"
	"reflect"
)

var (
	stringType      = reflect.TypeOf((*string)(nil)).Elem()
	stringSliceType = reflect.TypeOf([]string(nil))
)

func main() {
	ssv := reflect.MakeSlice(stringSliceType, 0, 10)

	sv := reflect.New(stringType).Elem()
	fmt.Println(sv.Kind())
	sv.SetString("abc")
	fmt.Println(sv)

	ssv = reflect.Append(ssv, sv)
	ss := ssv.Interface().([]string)
	fmt.Println(ss)

	var err error
	fmt.Println(err)
	err = nil
	fmt.Println(err == nil)
	err = errors.New("aaa")
	ivv := reflect.ValueOf(err)
	fmt.Println(ivv.IsValid(), ivv.Kind(), ivv.IsNil())
}

func hasNoValue(i interface{}) bool {
	iv := reflect.ValueOf(i)
	if !iv.IsValid() {
		return true
	}

	switch iv.Kind() {
	case reflect.Ptr, reflect.Slice, reflect.Map, reflect.Func, reflect.Interface:
		return iv.IsNil()
	default:
		return false
	}
}

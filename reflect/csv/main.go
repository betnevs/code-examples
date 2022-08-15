package main

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

type MyData struct {
	Name   string `csv:"name"`
	Age    int    `csv:"age"`
	HasPet bool   `csv:"has_pet"`
}

// Marshal maps all structs in a slice of structs to a slice of slice of
// strings.
func Marshal(v interface{}) ([][]string, error) {
	sliceVal := reflect.ValueOf(v)
	if sliceVal.Kind() != reflect.Slice {
		return nil, errors.New("must be a slice of structs")
	}

	structType := sliceVal.Type().Elem()
	if structType.Kind() != reflect.Struct {
		return nil, errors.New("must be a slice of structs")
	}

	var out [][]string
	header := marshalHeader(structType)
	out = append(out, header)

	for i := 0; i < sliceVal.Len(); i++ {
		row, err := marshalOne(sliceVal.Index(i))
		if err != nil {
			return nil, err
		}
		out = append(out, row)
	}
	return out, nil
}

func marshalOne(vv reflect.Value) ([]string, error) {
	var row []string

	vt := vv.Type()
	for i := 0; i < vv.NumField(); i++ {
		fieldValue := vv.Field(i)
		if _, ok := vt.Field(i).Tag.Lookup("csv"); !ok {
			continue
		}
		switch fieldValue.Kind() {
		case reflect.Int:
			row = append(row, strconv.FormatInt(fieldValue.Int(), 10))
		case reflect.String:
			row = append(row, fieldValue.String())
		case reflect.Bool:
			row = append(row, strconv.FormatBool(fieldValue.Bool()))
		default:
			return nil, fmt.Errorf("cannot handle field of kind %v", fieldValue.Kind())
		}
	}
	return row, nil
}

func marshalHeader(vt reflect.Type) []string {
	var row []string
	for i := 0; i < vt.NumField(); i++ {
		field := vt.Field(i)
		if curTag, ok := field.Tag.Lookup("csv"); ok {
			row = append(row, curTag)
		}
	}
	return row
}

// Unmarshal maps all of the rows of data in a slice of slice strings
// into a slice of structs.
func Unmarshal(data [][]string, v interface{}) error {

}

func main() {
	//datas := []MyData{
	//	{"aaa", 11, true},
	//	{"bb", 101, false},
	//}
	//Marshal(datas)
	test := MyData{
		HasPet: false,
		Age:    0,
		Name:   "ddddd",
	}

	t := reflect.TypeOf(test)

}

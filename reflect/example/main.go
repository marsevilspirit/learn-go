package main

import (
	"fmt"
	"reflect"
)

func FirstCopy(data []byte, i interface{}) error {
	i = &data
	return nil
}

func SecondCopy(data []byte, i interface{}) error {
	v := reflect.Indirect(reflect.ValueOf(i))
	v.SetBytes(data)
	return nil
}

type ByteSlice struct {
	Data []byte
}

func main() {
	//var first ByteSlice
	//var second ByteSlice

	first := []byte{0, 0, 0}
	second := []byte{0, 0, 0}

	FirstCopy([]byte{1, 2, 3}, &first)
	SecondCopy([]byte{1, 2, 3}, &second)

	fmt.Println(first)
	fmt.Println(second)
}

package main

import (
	"reflect"
	"fmt"
)

type NB struct {
	Name string
}

func main() {
	nb := NB{Name: "mars"}
	valueof := reflect.ValueOf(nb)
	typeof := reflect.TypeOf(nb)
	fmt.Println(valueof)
	fmt.Println(typeof)
}

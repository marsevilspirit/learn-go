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
	typeof := reflect.TypeOf(nb)
	fmt.Println("TypeOf:", typeof)
}

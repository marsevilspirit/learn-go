package main

import (
    "fmt"
)

type animal interface {
    speak() string
}

type dog struct {
    name string
}

func (d dog) speak() string {
    return d.name + " Woof!"
}

type cat struct {
    name string
}

func (c cat) speak() string {
    return c.name + " Meow!"
}

func main() {
    d := dog{"Rover"}
    c := cat{"Whiskers"}

    animals := []animal{d, c}

    for _, a := range animals {
        fmt.Println(a.speak())
    }
}

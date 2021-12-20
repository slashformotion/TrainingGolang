package main

import (
	"fmt"
)

type Person struct {
	nam string
	age int
}

type Cooker interface {
	Cook()
}
type Chef struct {
	Person
	Stars int
}

func (c Chef) Cook() {
	fmt.Printf("I'm cooking with %v stars\n", c.Stars)
}

func processPerson(i interface{}) {
	switch p := i.(type) {
	case Person:
		fmt.Println("person")
	case Chef:
		fmt.Println("chef")
		p.Cook()
	default:
		fmt.Println("Not a person")
	}
}

func main() {
	var x interface{} = Chef{
		Person: Person{"bob", 10},
		Stars:  20,
	}
	fmt.Printf("Type: %T, value: %v\n", x, x)
	s, ok := x.(string)
	fmt.Printf("Person as string %v, value: %v\n", ok, s)
	processPerson(x)
}

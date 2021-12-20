package main

import "fmt"

type Animal interface {
	Speak() string
}

type Dog struct {
	color string
}

func (d Dog) Speak() string {
	return "Woof"
}

type Cat struct {
	jumpHeight int
}

func (c Cat) Speak() string {
	return "Miaou"
}

func main() {
	animals := []Animal{
		Dog{"white"},
		Cat{2},
	}
	for _, a := range animals {
		fmt.Println(a.Speak())
		fmt.Printf("Type of animal %T\n", a)
		// Type assertion

		if t, ok := a.(Dog); ok {
			// a is a dog
			fmt.Printf("Type of animal %v, %v\n", t, ok)
			fmt.Printf("we a have a dog with color%v\n", t.color)
		} else {
			fmt.Printf("Type of animal %v, %v\n", t, ok)
			fmt.Printf("we a have a cat with jumpHeight\n")

		}
	}
	fmt.Println("-------------------------------")
	for _, a := range animals {
		describeAnimal(a)
	}
}

func describeAnimal(a Animal) {
	switch v := a.(type) { // v prend la valeur de a avec le bon type
	case Dog:
		fmt.Printf("we a have a dog with color%v\n", v.color)
	case Cat:
		fmt.Printf("we a have a cat with jumpHeight: %v\n", v.jumpHeight)
	}
}

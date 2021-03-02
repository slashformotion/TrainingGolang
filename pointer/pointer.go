//pointer.o
package main

import (
	"fmt"
)

func updateVal(ptr *string, newVal string) {
	*ptr = newVal
}

func main() {
	fmt.Println("--- Training.go => pointer ---")

	i := 1
	var p *int = &i
	fmt.Printf("i=%v\n", i)
	fmt.Printf("p=%v\n", p)
	fmt.Printf("*p=%v\n", *p)
	fmt.Println("------------")

	s := "Bob"
	sPtr := &s
	s2 := *sPtr
	fmt.Println("String pointer")
	fmt.Printf("s=%v\n", s)
	fmt.Printf("sPtr=%v\n", sPtr)
	fmt.Printf("*sPtr=%v\n", *sPtr)
	fmt.Printf("s2=%v\n", s2)
	fmt.Println("------------")
	fmt.Println("Modification de s à \"truc\"")
	s = "truc"
	fmt.Printf("s=%v\n", s)
	fmt.Printf("sPtr=%v\n", sPtr)
	fmt.Printf("*sPtr=%v\n", *sPtr)
	fmt.Printf("s2=%v\n", s2)
	fmt.Println("------------")
	*sPtr = "Quel plaisir"
	fmt.Printf("s=%v\n", s)
	fmt.Printf("sPtr=%v\n", sPtr)
	fmt.Printf("*sPtr=%v\n", *sPtr)
	fmt.Println("------------")
	fmt.Println("Utilisation de la fonction avec pointer updateVal")
	updateVal(&s, "incroyable")
	fmt.Printf("s=%v\n", s)
	fmt.Printf("sPtr=%v\n", sPtr)
	fmt.Printf("*sPtr=%v\n", *sPtr)
	fmt.Println("------------")

}

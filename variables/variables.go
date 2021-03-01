package main

import "fmt"

var pi float32 = 3.14

func main() {
	var age uint8 = 20
	var age2 uint8 // ona int16, int8, byte, ect...

	fmt.Println(age, age2)

	var weight, height int16 = 80, 190

	fmt.Println(height, weight)

	var name = "Bobe"
	email := "bob@go.org"

	fmt.Println(name, email)
}

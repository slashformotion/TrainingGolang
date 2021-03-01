package main

import "fmt"

// Greetings ...
func Greetings() {

	fmt.Println("Hello")
}

// OldEnough ...
func OldEnough(age int) bool {
	return age > 18
}

func printInfo(name string, age int, email string) {
	fmt.Printf("Name=%s, age=%d, email='%s' \n", name, age, email)
}

func avg(x, y float64) float64 {
	return (x + y) / 2
}

func sumNamedReturn(x, y, z int) (sum int) {
	sum = x + y + z // ici on se rend compte que 'sum' existe déjà dans l'espace local de la fonction
	return sum
}

func main() {
	fmt.Println(OldEnough(12))
	printInfo("bobi", 15, "bidule@email.com")
	var result float64 = avg(12, 14)
	fmt.Println(result)

}

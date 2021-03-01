package main

import "fmt"

func main() {
	var tabl [3]string
	fmt.Printf("name=%v (len=%v)\n", tabl, len(tabl))
	tabl[0] = "Bob"
	tabl[2] = "Bobi"

	fmt.Printf("name=%v (len=%v)\n", tabl, len(tabl))

	odds := [4]int{1, 3, 5, 7}
	fmt.Printf("odds=%v (len=%v)\n", odds, len(odds))

}

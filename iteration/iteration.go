package main

import (
	"fmt"
)

func main() {
	for i := 0; i <= 5; i++ {
		fmt.Println(i)
	}

	i := 0
	for i <= 5 {
		fmt.Println(i)
		i++
	}

	i = 0
	for {
		i++
		fmt.Printf("i = %d\n", i)

		if i%2 != 0 {
			fmt.Println("Odd looping")
			continue
		} else {
			fmt.Println("Even looping")

		}
		if i >= 11 {
			fmt.Println("Loop end")
			break
		}

	}

}

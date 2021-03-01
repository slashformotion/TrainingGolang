//range.o
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Training.go")

	tabl := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i, num := range tabl {
		fmt.Printf("nums[%d]=%d\n", i, num)
	}

	var testStr string = "BOBIBOUB"
	for i, chare := range testStr {
		fmt.Printf("nums[%d]=%c\n", i, chare)

	}
}

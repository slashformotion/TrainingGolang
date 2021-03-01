package main

import "fmt"

func tir() {
	fmt.Println("--------------")
}

func main() {
	// Initialisation Slices
	nums := make([]int, 2, 3)
	nums[0] = 10
	nums[1] = -1
	fmt.Printf("%v (len=%d, cap=%d), type=%T \n", nums, len(nums), cap(nums), nums)

	nums = append(nums, 64)
	fmt.Printf("%v (len=%d, cap=%d), type=%T \n", nums, len(nums), cap(nums), nums)

	nums = append(nums, 38)
	fmt.Printf("%v (len=%d, cap=%d), type=%T \n", nums, len(nums), cap(nums), nums)

	tir()

	// SUBSLICES
	letters := []string{"g", "o", "l", "a", "n", "g"}
	fmt.Printf("%v (len=%d, cap=%d), type=%T \n", letters, len(letters), cap(letters), letters)

	sub1 := letters[0:2] // attention: ici le 3eme elem n'est pas dans le tableau
	sub2 := letters[2:]
	fmt.Printf("%v (len=%d, cap=%d), type=%T \n", sub1, len(sub1), cap(sub1), sub1)
	fmt.Printf("%v (len=%d, cap=%d), type=%T \n", sub2, len(sub2), cap(sub2), sub2)
	tir()
	sub2[0] = "UP"

	fmt.Printf("%v (len=%d, cap=%d), type=%T \n", letters, len(letters), cap(letters), letters)

	fmt.Printf("%v (len=%d, cap=%d), type=%T \n", sub1, len(sub1), cap(sub1), sub1)

	fmt.Printf("%v (len=%d, cap=%d), type=%T \n", sub2, len(sub2), cap(sub2), sub2)

	// Copie de slices
	tir()

	subCopy := make([]string, len(sub1))
	copy(subCopy, sub1)
	subCopy[0] = "Down"

	fmt.Printf("%v (len=%d, cap=%d), type=%T \n", subCopy, len(subCopy), cap(subCopy), subCopy)
	fmt.Printf("%v (len=%d, cap=%d), type=%T \n", letters, len(letters), cap(letters), letters)
	fmt.Printf("%v (len=%d, cap=%d), type=%T \n", sub1, len(sub1), cap(sub1), sub1)

}

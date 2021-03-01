//errors.o
package main

import (
	"errors"
	"fmt"
	"io/ioutil"
)

func myFunc(condition bool) (int, error) {
	if !condition {
		return 0, errors.New("Hello error")
	}
	return 0, nil
}

func readFile(filename string) (string, error) {
	dat, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}
	if len(dat) == 0 {
		return "", fmt.Errorf("Empty content (filename=%s)", filename)
	}
	return string(dat), nil
}

func main() {
	fmt.Println("Training.go => errors")
	value, err := readFile("dumb.txt")
	if err != nil {
		fmt.Printf("Error while reading file: %v\n", err)
		return
	}
	fmt.Printf("value=\"%s\"\n", value)
}

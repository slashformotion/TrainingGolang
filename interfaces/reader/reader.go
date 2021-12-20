package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	// r := strings.NewReader("Hello gphers ! Readers are awe some"),
	r, err := os.Open("txt.txt")
	if err != nil {
		fmt.Printf("can't open file: %v\n", err)
		return
	}
	buff, err := ioutil.ReadAll(r)
	if err != nil {
		fmt.Printf("Error in Reader: %v\n", err)
		return
	}
	fmt.Println(string(buff))

	resp, err := http.Get("https://golang.org")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	buffs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error in Reader: %v\n", err)
		return
	}
	fmt.Println(string(buffs))
}

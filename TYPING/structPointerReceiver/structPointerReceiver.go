//structPointerReceiver.o
package main

import (
	"fmt"
	"strings"
)

type User struct {
	Name string
	Age  uint8
}

func (u User) IsMajeur() bool {
	return u.Age >= 18
}

func (u *User) toLower() {
	u.Name = strings.ToLower(u.Name)
}

func main() {
	fmt.Println("--- Training.go => structPointerReceiver ---")
	u := User{
		Name: "BOBINETTE",
		Age:  22,
	}
	fmt.Printf("%v\n", u)
	u.toLower()
	fmt.Printf("%v\n", u)

	incUser := &User{
		Name: "ROBINETTE",
		Age:  25,
	}

}

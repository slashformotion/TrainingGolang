package main

import "fmt"

// User ...
type User struct {
	Name  string
	Email string
}

// Admin ...
type Admin struct {
	User
	Level int
}

func main() {
	u1 := User{
		Name:  "Bob",
		Email: "bob@email.com",
	}
	fmt.Printf("User 1: %v\n", u1)
	a := Admin{
		Level: 2,
	}
	a.Name = "Bobi"
	a.Email = "bobi@email.com"
	fmt.Printf("Admin 1: %v\n", a)

	a2 := Admin{
		Level: 2,
		User: User{
			Name:  "Bidule",
			Email: "bidule@email.com",
		},
	}
	fmt.Printf("Admin 2: %v\n", a2)

}

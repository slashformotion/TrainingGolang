package main

import (
	"fmt"

	"./player"
)

func main() {
	var p1 player.Player

	fmt.Printf("Player 1: %v\n", p1)
	fmt.Printf("Player 1: name=%v , age=%d\n", p1.Name, p1.Age)
	a := player.Avatar{"http://avatar.jpg"}
	p1.Avatar = a
	fmt.Printf("Avatar : %v\n", a)

	p3 := player.Player{
		Name: "John",
		Age:  25,
		Avatar: player.Avatar{
			Url: "http://url.com",
		},
	}
	fmt.Printf("Player 3 : %v\n", p3)

	p4 := player.New("Bobette")
	p4.Avatar = a

	fmt.Printf("Player 4 : %v\n", p4)

	p4.Avatar.Url = "truc"
	fmt.Printf("Player 4 : %v\n", p4)

	fmt.Printf("Player 1: %v\n", p1)

}

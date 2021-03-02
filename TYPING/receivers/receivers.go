package main

import "fmt"

// Rect struc: represent a rectangle
type Rect struct {
	Width, Height int
}

// Area return area
func (r Rect) Area() int {
	return r.Width * r.Height
}

func main() {
	r := Rect{2, 4}
	fmt.Printf("Rect area=%d\n", r.Area())
	fmt.Println(r)
}

// ATTENTION ici on a que des "value receiver",
// en clair on ne travaille pas sur des objets mais sur une copie de la structure

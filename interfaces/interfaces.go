//interfaces.o
package main

import (
	"fmt"
)

type Saver interface {
}

type Loader interface {
}

type SaverLoader interface {
	Saver
	Loader
}

type Instrument interface {
	Play()
}

type Guitar struct {
	Notes int
}

func (g *Guitar) Play() {
	fmt.Printf("len(notes): %v\n", g.Notes)
}

func main() {
	fmt.Println("--- Training.go => interfaces ---")
	var instru Instrument
	instru = &Guitar{1}
	instru.Play()
}

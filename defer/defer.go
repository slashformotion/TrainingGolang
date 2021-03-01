//defer.o
package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("--Training.go => defer--")
	var filename string = "dumb.txt"
	f, err := os.OpenFile(filename, os.O_RDONLY, 0755)
	defer f.Close() // sera exécuté quand nous sortirons de main
	if err != nil {
		log.Fatalf("Can't open file (filename=%s) : %v", filename, err)
	}
}

// le defer fonctionne en mode LIFO = Last In First Outs

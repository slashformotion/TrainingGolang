package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/slashformotion/TrainingGolang/PROJECTS/dictionary/dictionary"
)

func main() {
	d, err := dictionary.New("./badger")
	HandleErr(err)
	defer d.Close()
	action := flag.String("action", "list", "action to perform on the dictionary")
	flag.Parse()
	switch *action {
	case "list":
		actionList(d)
	case "add":
		actionAdd(d, flag.Args())
	case "define":
		actionDefine(d, flag.Args())
	case "remove":
		actionRemove(d, flag.Args())
	default:
		fmt.Printf("Unknown action: %v\n", *action)
	}
}

func actionList(d *dictionary.Dictionary) {
	words, entries, err := d.List()
	HandleErr(err)
	for _, word := range words {
		fmt.Println(entries[word])
	}
}
func actionAdd(d *dictionary.Dictionary, args []string) {
	word := args[0]
	def := args[1]
	err := d.Add(word, def)
	HandleErr(err)
	fmt.Printf("'%v' added to the dictionary\n", word)
}

func actionDefine(d *dictionary.Dictionary, args []string) {
	var word string = args[0]
	e, err := d.Get(word)
	HandleErr(err)
	fmt.Printf("%v\n", e)
}

func actionRemove(d *dictionary.Dictionary, args []string) {
	var word string = args[0]
	err := d.Remove(word)
	HandleErr(err)
	fmt.Printf("'%v' has been deleted\n", word)
}
func HandleErr(err error) {
	if err != nil {
		fmt.Printf("Dictionary error: %v\n", err)
		os.Exit(1)
	}
}

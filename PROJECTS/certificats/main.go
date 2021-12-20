package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/slashformotion/TrainingGolang/PROJECTS/certificats/cert"
	"github.com/slashformotion/TrainingGolang/PROJECTS/certificats/html"
	"github.com/slashformotion/TrainingGolang/PROJECTS/certificats/pdf"
)

func main() {
	outputType := flag.String("type", "pdf", "ouput type of the certificate.")
	file := flag.String("file", "students.csv", "the file to save")
	flag.Parse()

	var saver cert.Saver
	var err error
	switch *outputType {
	case "html":
		saver, err = html.New("output")
	case "pdf":
		saver, err = pdf.New("output")
	default:
		fmt.Printf("Unknown output format: go=%v\n", *outputType)
	}
	if err != nil {
		fmt.Printf("Could not create egenrator: %v\n", err)
		os.Exit(1)
	}
	if err != nil {
		fmt.Printf("Error during certificate creation: %v", err)
		os.Exit(1)
	}
	cs, err := cert.ParseCsv(*file)
	if err != nil {
		fmt.Printf("Error during csv parsing: %v", err)
		os.Exit(1)
	}
	for _, c := range cs {
		saver.Save(*c)
	}

}

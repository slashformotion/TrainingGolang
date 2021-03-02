package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// FindReplaceFile ...
func FindReplaceFile(src, dst, old, new string) (occ int, lines []int, err error) {
	// Open src file
	srcFile, err := os.Open(src)
	if err != nil {
		return 0, nil, err
	}
	defer srcFile.Close()

	//writing the new file
	dstFile, err := os.Create(dst)
	if err != nil {
		return 0, nil, err
	}
	defer dstFile.Close()

	// prepare processing
	old += " "
	new += " "
	var i int = 1

	scanner := bufio.NewScanner(srcFile)
	writer := bufio.NewWriter(dstFile)
	defer writer.Flush()

	for scanner.Scan() {
		found, res, occLine := ProcessLine(scanner.Text(), old, new)
		if found {
			lines = append(lines, i)
			occ += occLine
		}
		fmt.Fprintln(writer, res)
		i++

	}
	return occ, lines, nil
}

// ProcessLine ..
func ProcessLine(line, old, new string) (found bool, res string, occ int) {
	if !strings.Contains(line, old) {
		return false, line, 0
	}
	found = true
	occ = strings.Count(line, old)
	res = strings.ReplaceAll(line, old, new)
	return found, res, occ
}

func main() {
	var src, dst, old, new string = "dumb.txt", "dumber.txt", "is", "bite"
	occ, lines, err := FindReplaceFile(src, dst, old, new)
	if err != nil {
		log.Fatalf("An error happened: %v", err)
		return
	}
	fmt.Println("== Summary ==")
	fmt.Printf("Number of occurence of %s: %d\n", old, occ)
	fmt.Printf("Number of lines: %d\n", len(lines))
	fmt.Printf("Lines: %v\n", lines)

	fmt.Println("== End of Summary ==")

}

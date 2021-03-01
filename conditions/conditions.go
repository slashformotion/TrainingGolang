package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Give me your age : ")
	ageStr, _err := reader.ReadString('\n')
	ageStr = strings.Replace(ageStr, "\n", "", 1)
	fmt.Println("votre age est : ", ageStr)

	if _err != nil {
		fmt.Println(_err.Error())
	}
	var age, _ = strconv.Atoi(ageStr)
	fmt.Println("votre age est : ", age)

	if age < 18 {
		fmt.Println("Vous avez moins de 18 ans, vous etes mineurs")
	} else if age <= 60 {
		fmt.Println("Vous avez moins de 60 ans, vous n'etes pas encore à la retraite.")
	} else {
		fmt.Println("Alors la retraite c'est comment ? ")
	}

	switch age {
	case 30:
		fmt.Println("Crise de la 50aine incomming")
	case 60:
		fmt.Println("Crise de la retraite.")
	}
}

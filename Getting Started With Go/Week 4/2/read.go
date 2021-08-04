package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	type name struct {
		fname string
		lname string
	}

	var names []name
	var filename string

	fmt.Print("Enter filename, you don't need to add extention : ")
	fmt.Scanln(&filename)

	filenameWithExtension := filename + ".txt"

	file, err := os.Open(filenameWithExtension)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		var n name
		s := strings.Fields(scanner.Text())
		n.fname = s[0]
		n.lname = s[1]

		names = append(names, n)
	}

	file.Close()

	for _, nm := range names {
		fmt.Printf("First Name : %s, Last Name : %s \n", nm.fname, nm.lname)
	}
}

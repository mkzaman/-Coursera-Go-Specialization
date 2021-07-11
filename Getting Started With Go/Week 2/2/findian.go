package main

import (
	"fmt"
	"strings"
)

func main() {
	var temp string
	fmt.Scanln(&temp)
	result := strings.Contains(temp, "a & i & n")

	if result {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}

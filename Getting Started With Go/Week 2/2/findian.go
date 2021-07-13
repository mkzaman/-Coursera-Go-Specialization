package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	temp := scanner.Text()
	caseInsensitive := strings.ToLower(temp)
	first := string(caseInsensitive[0])
	last := string(caseInsensitive[len(caseInsensitive)-1:])

	if strings.Compare(first, "i") == 0 && strings.Compare(last, "n") == 0 && strings.Contains(caseInsensitive, "a") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}

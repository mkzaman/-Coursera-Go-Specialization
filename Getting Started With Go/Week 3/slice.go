package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {

	var s []int
	s = make([]int, 0, 3)

	for {
		var input string
		fmt.Scanln(&input)
		if strings.Compare(input, "X") == 0 {
			break
		}
		intVal, error := strconv.Atoi(input)

		if error != nil {
			fmt.Println(error)
		}
		s = append(s, intVal)
		sort.Ints(s)
		fmt.Println(s)
	}
}

package main

import "fmt"

func main() {
	var real float64
	var truncated int64
	fmt.Scanln(&real)
	truncated = int64(real)
	fmt.Println(truncated)
}

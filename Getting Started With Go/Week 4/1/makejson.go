package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var m = make(map[string]string)
	var name string
	var address string
	fmt.Print("Enter Name : ")
	fmt.Scanln(&name)
	fmt.Print("Enter Address : ")
	fmt.Scanln(&address)

	m["name"] = name
	m["addrress"] = address

	jsn, err := json.Marshal(m)

	if err != nil {
		fmt.Println("There was a error")
	} else {
		fmt.Println(string(jsn))
	}
}

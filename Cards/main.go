package main

import "fmt"

func main() {
	var strname string
	fmt.Println("Please enter 2 Letter state name")
	fmt.Scanln(&strname)
	fmt.Println(stateName(strname))
}

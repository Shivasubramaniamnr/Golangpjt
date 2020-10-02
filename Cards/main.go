package main

import "fmt"

func main() {
	var strname string

	fmt.Println("Please enter 2 Letter state name")

	fmt.Scanln(&strname)

	// stateName func is to find the abbreviation of 2 letter state code
	fmt.Println(stateName(strname))
}

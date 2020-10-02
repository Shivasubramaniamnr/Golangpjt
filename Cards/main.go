package main

import "fmt"

func main() {
	var strname string

	for k := 0; k <= 10; k++ {
		fmt.Println("Please enter 2 Letter state name, to stop enter STOP")

		fmt.Scanln(&strname)

		if strname == "STOP" {
			break
		}
		// stateName func is to find the abbreviation of 2 letter state code
		fmt.Println(cityName(strname))
	}
}

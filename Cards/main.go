package main

import "fmt"

func main() {
	var stname string
	var ctname string

	for k := 0; k <= 10; k++ {
		fmt.Println("Please enter 2 Letter state name and city name , to stop enter STOP")

		fmt.Scanln(&stname)
		fmt.Scanln(&ctname)

		if stname == "STOP" {
			break
		}
		// stateName func is to find the abbreviation of 2 letter state code
		fmt.Println(cityName(stname, ctname))
	}
}

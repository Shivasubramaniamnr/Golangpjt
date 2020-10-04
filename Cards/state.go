package main

import (
	"fmt"
	"sort"
	"strings"
)

var cityNames []string
var cityString string
var pos int

func cityName(state string, city string) string {
	cityNames = nil

	switch state {
	case "TN":
		cityNames = append(cityNames, "Chennai", "Madurai", "Trichy", "Nellai", "Coimbatore", "Salem")
	case "AP":
		cityNames = append(cityNames, "Vizag", "Guntur", "Nellore", "Vijaywada")
	case "DL":
		cityNames = append(cityNames, "Rohini", "PitamPura", "Karol Bagh", "CP")
	default:
		cityNames = append(cityNames, "No information about cities found")
	}
	sort.Strings(cityNames)
	pos = sort.StringSlice(cityNames).Search(city)
	fmt.Println(pos)

	if pos > len(cityNames) {
		cityString = "City does not belong to the State"
	} else {
		cityString = strings.Join(cityNames, ",")
	}

	return cityString
}

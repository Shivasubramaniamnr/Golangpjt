package main

import (
	"sort"
	"strings"
)

var cityNames []string
var cityString string
var pos int

func cityName(state string, city string) string {
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
	pos = sort.SearchStrings(cityNames, city)
	if pos > 0 {
		cityString = strings.Join(cityNames, ",")
	} else {
		cityString = "City does not belong to the State"
	}

	return cityString
}

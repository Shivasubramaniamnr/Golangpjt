package main

import (
	"sort"
)

var cityNames []string
var cityString string
var pos int

func cityName(state string, city string) string {
	cityNames = nil
	pos = 0

	switch state {
	case "TN":
		cityNames = []string{"Chennai", "Madurai", "Trichy", "Nellai", "Coimbatore", "Salem"}
	case "AP":
		cityNames = []string{"Vizag", "Guntur", "Nellore", "Vijaywada"}
	case "DL":
		cityNames = []string{"Rohini", "PitamPura", "Karol Bagh", "CP"}
	}
	sort.Strings(cityNames)

	pos = sort.SearchStrings(cityNames, city)

	if city == cityNames[pos] {
		cityString = "City is found in the state"
	} else {
		cityString = "City does not belong to the State"
	}

	return cityString
}

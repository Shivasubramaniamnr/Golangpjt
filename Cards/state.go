package main

import "strings"

var cityNames []string
var cityString string

func cityName(state string) string {
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
	cityString = strings.Join(cityNames, ",")
	return cityString
}

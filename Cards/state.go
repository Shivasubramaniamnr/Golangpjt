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
		cityNames = []string{"Chennai", "Madurai", "Trichy", "Nellai", "Coimbatore", "Salem"}
	case "AP":
		cityNames = []string{"Vizag", "Guntur", "Nellore", "Vijaywada"}
	case "DL":
		cityNames = []string{"Rohini", "PitamPura", "Karol Bagh", "CP"}
	}
	sort.Strings(cityNames)

	pos = sort.SearchStrings(cityNames, city)
	fmt.Println(cityNames)
	fmt.Println(pos)

	if pos <= len(cityNames) {
		cityString = strings.Join(cityNames, ",")
	} else {
		cityString = "City does not belong to the State"
	}

	return cityString
}

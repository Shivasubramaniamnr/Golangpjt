package main

var statename string

func sName(state string) string {
	if state == "TN" {
		statename = "Tamil Nadu"
	} else {
		statename = "State Abbreviation not found"
	}
	return statename
}

func stateName(state string) string {
	switch state {
	case "TN":
		statename = "Tamil Nadu"
	case "AP":
		statename = "Andhra Pradesh"
	case "DL":
		statename = "Delhi"
	}
	return statename
}

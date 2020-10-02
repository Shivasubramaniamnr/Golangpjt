package main

var statename string

func stateName(state string) string {
	switch state {
	case "TN":
		statename = "Tamil Nadu"
	case "AP":
		statename = "Andhra Pradesh"
	case "DL":
		statename = "Delhi"
	default:
		statename = "State Abbreviation not found"
	}
	return statename
}

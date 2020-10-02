package main

import "fmt"
var statename string

func sName (string state) string{
if state == "TN"
   return ("Tamil Nadu")
}

func stateName (string state) string{
	switch state {
	case "TN": statename = "Tamil Nadu"
	case "AP": statename = "Andhra Pradesh"
	case "DL": statename = "Delhi"
	}
	return (statename)
}
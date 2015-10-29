package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	miTokm = 1.60934
	miTometer = 1609.34
	miTofeet = 5280

	kmTometer = 1000
	kmTomi = 0.621371
	kmTofeet = 3280.84

	meterTofeet = 3.28084
	meterTokm = 0.001
	meterTomi = 0.000621371

	ftTomi = 0.000189394
	ftTokm = 0.0003048
	ftTometer = 0.3048
)

//this is from Go Day 2 Part 6
func main() {
	from := os.Args[1]
	to := os.Args[2]

	switch {
	case strings.HasSuffix(from, "mi"):
		fmt.Println("From Miles")
		switch to {
		case "km":
			//this is like a substring
			//we get the length of the string and go back 2 spaces
			miles, _ := strconv.ParseFloat(from[0:len(from)-2], 64)
			fmt.Println(miles * miTokm)
		case "m":
			miles, _ := strconv.ParseFloat(from[0:len(from)-2], 64)
			fmt.Println(miles * miTometer)
		case "ft":
			miles, _ := strconv.ParseFloat(from[0:len(from)-2], 64)
			fmt.Println(miles * miTofeet)
		}
	case strings.HasSuffix(from, "km"):
		fmt.Println("From Kilometers")
		switch to {
		case "mi":
			kilometer, _ := strconv.ParseFloat(from[0:len(from)-2], 64)
			fmt.Println(kilometer * kmTomi)
		case "ft":
			kilometer, _ := strconv.ParseFloat(from[0:len(from)-2], 64)
			fmt.Println(kilometer * kmTofeet)
		case "m":
			kilometer, _ := strconv.ParseFloat(from[0:len(from)-2], 64)
			fmt.Println(kilometer * kmTometer)
		}
	case strings.HasSuffix(from, "ft"):
		fmt.Println("From Feet")
		switch to {
		case "mi":
			feet, _ := strconv.ParseFloat(from[0:len(from)-2], 64)
			fmt.Println(feet * ftTomi)
		}
	}

}


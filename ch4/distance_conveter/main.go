package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	miTokm = 1.60934
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
		}
	}

}

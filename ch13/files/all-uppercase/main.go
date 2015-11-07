package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"log"
	"strings"
)

func main() {

	file, err := os.Open("upperme")

	if err != nil {
		log.Fatalln("There is no file")
	}

	bytes, err := ioutil.ReadAll(file)

	if err != nil {
		log.Fatalln("There is a problem Jim")
	}

	fmt.Println(bytes)
	fmt.Println(string(bytes))

	upperString := strings.ToUpper(string(bytes))

	fmt.Println(upperString)

}
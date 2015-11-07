package  main

import (
	"fmt"
	"os"
	"io/ioutil"
	"log"
	//"strings"
)

func main() {
	file, err := os.Open("cat.txt")

	if err != nil {
		log.Fatalln("No file here")
	}

	fmt.Println(file)
	bytes, err := ioutil.ReadAll(file)

	if err != nil {
		log.Fatalln("Something went wrong")
	}
	fmt.Println(bytes)
	fmt.Println(string(bytes))
}

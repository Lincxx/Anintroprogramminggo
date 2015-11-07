package main

import (
	"os"
	"log"
//	main
//"strings"
)

func main() {

	f, err := os.Create("test.txt")

	if err != nil {
		log.Fatalln("It's broken")
	}

	defer f.Close()

	str := "txt"
	bs := []byte(str)

	_, err = f.Write(bs)

	if err != nil {
		log.Fatalln("Error to file")
	}
}


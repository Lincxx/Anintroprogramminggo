package main

import (
	"os"
	"log"
//	main
//"strings"
)

func main() {

	err := os.Mkdir("somedir", 0x777)
	if err != nil {
		log.Fatalln("failed to creat dir")
	}

	f, err := os.Create("somedir/test.txt")

	if err != nil {
		log.Fatalln("It's broken")
	}

	defer f.Close()

	str := "txt"
	bs := []byte(str)

	//we are using =, because we already declared := on line 12
	_, err = f.Write(bs)

	if err != nil {
		log.Fatalln("Error to file")
	}
}


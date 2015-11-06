package main

import (
	"os"
	"log"
	"io/ioutil"
	"fmt"
	"strings"
)

func main() {
	f, err := os.Open("hello.txt")

	if err != nil {
		log.Fatalln("It's broken")
	}

	defer f.Close()

	//reader := strings.NewReader("some thing")

	bs, err := ioutil.ReadAll(f)

	if err != nil {
		log.Fatalln("Can't read")
	}

	str := string(bs)

	fmt.Println(bs)
	fmt.Println(str)

}

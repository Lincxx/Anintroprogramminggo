package main

import (
	"os"
	"log"
	"io/ioutil"
	"fmt"
	//"strings"
)

func main() {

//	fmt.Fprintln(os.Stderr,"Hello and Welcome. \nWhat is your name")
//	var name string
//	fmt.Scan(&name)


	f, err := os.Open(os.Args[1])

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

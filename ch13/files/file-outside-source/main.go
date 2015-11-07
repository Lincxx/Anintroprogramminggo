package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	//"strings"
	"net/http"
)

func main() {
	resp, err := http.Get("http://www.gutenberg.org/files/2781/old/jusss10.txt")

	if err != nil {
		log.Fatalln("It's broken")
	}
	file := resp.Write(resp)

	f, err := os.Open(file)

	if err != nil {
		log.Fatalln("It's broken")
	}

	//defer f.Close()

	//reader := strings.NewReader("some thing")

	bs, err := ioutil.ReadAll(f)

	if err != nil {
		log.Fatalln("Can't read")
	}

	str := string(bs)

	fmt.Println(bs)
	fmt.Println(str)

}

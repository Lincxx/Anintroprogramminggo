package main


import (
	"fmt"
	"os"
	"log"
	"bufio"
	"strings"
)
//This could be used to interact with a file line by line
func main() {
	srcFile, err := os.Open("hello.txt")
	if err != nil {
		log.Fatalln(err)
	}

	defer srcFile.Close()

	scanner := bufio.NewScanner(srcFile)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(">>>", strings.ToUpper(string(line[0])) + strings.ToLower(line[1:]))
	}
}

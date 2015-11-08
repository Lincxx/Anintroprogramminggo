package main

import (
	"fmt"
	"io"
	"log"
	"os"
	//"strings"
)

func cp(srcFileName, dstFileName string) error {
	srcF, err := os.Open(srcFileName)

	if err != nil {
		return fmt.Errorf("Can't open file: %v", err)
	}

	defer srcF.Close()

	dstF, err := os.Create(dstFileName)

	if err != nil {
		return fmt.Errorf("It's broken: %v")
	}

	defer dstF.Close()

	_, err = io.Copy(dstF, srcF)
	if err != nil {
		return fmt.Errorf("error writing destination file %v")
	}

	return nil
}

func main() {

	srcFileName := os.Args[1]
	dstFileName := os.Args[2]

	err := cp(srcFileName, dstFileName)
	if	err != nil {
		log.Fatalln(err)
	}

}

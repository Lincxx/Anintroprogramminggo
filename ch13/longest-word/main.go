package main
import (
	"fmt"
//"strings"
	"os"
	"log"
	"io"
	"bufio"
	//"strings"
)

func LongestWord(rdr io.Reader) string{

	scanner := bufio.NewScanner(rdr)
	scanner.Split(bufio.ScanWords)
	//tmpHolder := 0
	bigWord := ""
	for scanner.Scan() {
		word := scanner.Text()
		if len(word) > len(bigWord) {
			bigWord = word
		}
	}

	return bigWord
}

func main() {

	srcF, err := os.Open("moby.txt")

	if err != nil {
		log.Fatalln("Can't open file: %v", err)
	}

	longest := LongestWord(srcF)
	fmt.Println("Number of whales: ",longest)
}

//looks for a certain word in files result["word"]
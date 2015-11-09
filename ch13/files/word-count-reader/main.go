package main
import (
	"fmt"
	//"strings"
	"os"
	"log"
	"io"
	"bufio"
	"strings"
)

func WordCount(rdr io.Reader) map[string]int{

	counts := map[string]int{}
	scanner := bufio.NewScanner(rdr)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		word := scanner.Text()
		word = strings.ToUpper(word)
		counts[word]++
	}

	return counts
}

func main() {

	srcF, err := os.Open("moby.txt")

	if err != nil {
		log.Fatalln("Can't open file: %v", err)
	}

	counts := WordCount(srcF)
	fmt.Println("Number of whales: ",counts["whale"])
}

//looks for a certain word in files result["word"]
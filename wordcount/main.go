package main
import (
	"fmt"
	"strings"
)

func WordCount(str string) map[string]int{
	count := map[string]int{}

	for _, word := range strings.Fields(str) {
		count[word]++
	}
	return count
}

func main() {
	str := "test test"
	result := WordCount(str)

	fmt.Println(result)
}
package main
import (
	"fmt"
	"strings"
)

func WordCount(str string) map[string]int{
	count := map[string]int{}

	for _, v := range strings.Fields(str) {
		count[v]++
	}
	return count
}

func main() {
	str := "I can make this whatever I want"
	result := WordCount(str)

	fmt.Println(len(result))
}
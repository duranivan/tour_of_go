package main

import (
	"fmt"
	"strings"

	"golang.org/x/tour/wc"
)

// WordCount returns a map of the counts of each "word" in the string s
func WordCount(s string) map[string]int {
	m := make(map[string]int)

	for _, w := range strings.Fields(s) {
		// if elem, ok := m[w]; ok == true {
		// 	m[w]++
		// } else {
		// 	m[w] = elem + 1
		// }
		m[w]++
	}

	return m
}

func main() {
	fmt.Println(WordCount("hello world! hello people"))
	wc.Test(WordCount)
}

package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	r := make(map[string]int)
	for _, w := range strings.Fields(s) {
		r[w] = r[w] + 1
	}

	return r
}

func main() {
	wc.Test(WordCount)
}

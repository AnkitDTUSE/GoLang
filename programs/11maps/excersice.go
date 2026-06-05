package main

import (
	"strings"
)

func WordCount(s string) map[string]int {
	count := make(map[string]int)
	words := strings.Fields(s)

	for _, word := range words {
		count[word]++
	}

	return count
}

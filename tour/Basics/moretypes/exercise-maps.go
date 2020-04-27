package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	sArray := strings.Fields(s)
	for _, v := range sArray {
		m[v]++
	}
	return m
}

func main() {
	wc.Test(WordCount)
}

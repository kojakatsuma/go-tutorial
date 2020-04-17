package main

import "fmt"

func pringSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	pringSlice(s)

	s = s[:0]
	pringSlice(s)

	s = s[:4]
	pringSlice(s)

	s = s[2:]
	pringSlice(s)
}

package main

import "fmt"

func invoke() string {
	fmt.Println("invoke in world")
	return "world"
}

func main() {
	defer fmt.Println(invoke())
	fmt.Println("hello")
}
